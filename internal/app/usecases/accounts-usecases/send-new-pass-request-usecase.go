package accountsusecases

import (
	"fmt"
	"net/http"
	"time"

	"github.com/henrique998/go-auth/internal/app/contracts"
	"github.com/henrique998/go-auth/internal/app/entities"
	appErr "github.com/henrique998/go-auth/internal/app/errors"
	"github.com/henrique998/go-auth/internal/configs/logger"
	"github.com/henrique998/go-auth/internal/infra/utils"
)

type SendNewPassRequestUseCase struct {
	Repo          contracts.AccountsRepository
	VTRepo        contracts.VerificationTokensRepository
	EmailProvider contracts.EmailProvider
}

func (uc *SendNewPassRequestUseCase) Execute(email string) appErr.IAppError {
	logger.Info("Init SendNewPassRequest UseCase")

	account := uc.Repo.FindByEmail(email)

	if account == nil {
		return appErr.NewAppError("account does not exists.", http.StatusNotFound)
	}

	code, err := utils.GenerateToken(10)
	if err != nil {
		logger.Error("Error trying to generate code!", err)
		return appErr.NewAppError("internal server error", http.StatusInternalServerError)
	}

	message := fmt.Sprintf("Your verification code is: %s. Please use this code to complete your password reset process. The code will expire in 10 minutes.", code)
	expiresAt := time.Now().Add(10 * time.Minute)

	verificationCode := entities.NewVerificationToken(code, account.ID, expiresAt)

	uc.VTRepo.Create(*verificationCode)
	uc.EmailProvider.SendMail(email, "Password reset", message)

	return nil
}
