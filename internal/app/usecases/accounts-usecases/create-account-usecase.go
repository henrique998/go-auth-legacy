package accountsusecases

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/henrique998/go-auth/internal/app/contracts"
	"github.com/henrique998/go-auth/internal/app/entities"
	appError "github.com/henrique998/go-auth/internal/app/errors"
	"github.com/henrique998/go-auth/internal/app/request"
	"github.com/henrique998/go-auth/internal/configs/logger"
	"github.com/henrique998/go-auth/internal/infra/utils"
)

type CreateAccountUseCase struct {
	Repo          contracts.AccountsRepository
	VTRepo        contracts.VerificationTokensRepository
	EmailProvider contracts.EmailProvider
}

func (uc *CreateAccountUseCase) Execute(req request.CreateAccountRequest) appError.IAppError {
	logger.Info("Init CreateAccount UseCase")

	account, err := uc.Repo.FindByEmail(req.Email)
	if err != nil {
		return err
	}

	if account != nil {
		logger.Error("Error while validate account.", errors.New("account already exists"))
		return appError.NewAppError("account already exists", 400)
	}

	pass_hash, passErr := utils.HashPass(req.Pass)
	if passErr != nil {
		logger.Error("Error trying to hash password.", passErr)
		return appError.NewAppError("Internal server error", 500)
	}

	data := entities.NewAccount(req.Name, req.Email, pass_hash, req.Phone)

	err = uc.Repo.Create(*data)
	if err != nil {
		logger.Error("Error trying to create account.", errors.New(err.GetMessage()))
		return appError.NewAppError("Internal server error", 500)
	}

	tokenString, tokenErr := utils.GenerateToken(10)
	if tokenErr != nil {
		logger.Error("Error trying to generate token.", passErr)
		return appError.NewAppError("Internal server error", 500)
	}

	expiresAt := time.Now().Add(time.Hour * 2)

	verificationToken := entities.NewVerificationToken(tokenString, data.ID, expiresAt)

	err = uc.VTRepo.Create(*verificationToken)
	if err != nil {
		logger.Error("Error trying to create verification token.", errors.New(err.GetMessage()))
		return appError.NewAppError("Internal server error", 500)
	}

	appBaseUrl := os.Getenv("BASE_URL")
	verificationUrl := fmt.Sprintf("%saccounts/verify-email?token=%s", appBaseUrl, tokenString)

	body := fmt.Sprintf(`Olá, 
	
	Por favor, verifique seu endereço de e-mail clicando no link abaixo:
	
	%s
	
	Se você não se cadastrou em nosso site, ignore este e-mail.
	
	Obrigado!`, verificationUrl)

	emailErr := uc.EmailProvider.SendMail(req.Email, "Account verification.", body)
	if emailErr != nil {
		logger.Error("Error trying to send verification email.", emailErr)
		return appError.NewAppError("Internal server error", 500)
	}

	return nil
}
