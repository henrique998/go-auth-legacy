package accountsusecases

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/henrique998/go-auth/internal/app/contracts"
	"github.com/henrique998/go-auth/internal/app/entities"
	appError "github.com/henrique998/go-auth/internal/app/errors"
	"github.com/henrique998/go-auth/internal/configs/logger"
	"github.com/henrique998/go-auth/internal/infra/utils"
)

type Send2faCodeUseCase struct {
	Repo                  contracts.AccountsRepository
	VTRepo                contracts.VerificationCodesRepository
	TwoFactorAuthProvider contracts.TwoFactorAuthProvider
}

func (uc *Send2faCodeUseCase) Execute(accountId string) appError.IAppError {
	logger.Info("Init Send2FACode UseCase")

	account := uc.Repo.FindById(accountId)

	if account == nil {
		return appError.NewAppError("account does not exists", http.StatusNotFound)
	}

	if account.Phone == nil {
		return appError.NewAppError("account must have an phone number to complete 2fa proccess", http.StatusUnauthorized)
	}

	code, err := utils.GenerateCode(10)
	if err != nil {
		logger.Error("Error while generate 2fa code.", err)
		return appError.NewAppError("internal server error.", http.StatusInternalServerError)
	}

	message := fmt.Sprintf("Your verification code is: %s. Please use this code to complete your 2FA process. The code will expire in 10 minutes.", code)
	expiresAt := time.Now().Add(10 * time.Minute)
	fromNumber := os.Getenv("TWILIO_FROM_PHONE_NUMBER")

	verificationCode := entities.NewVerificationCode(code, accountId, expiresAt)

	uc.VTRepo.Create(*verificationCode)
	err = uc.TwoFactorAuthProvider.Send(fromNumber, *account.Phone, message)
	if err != nil {
		logger.Error("Error trying to send 2fa code!", err)
		return appError.NewAppError("internal server error.", http.StatusInternalServerError)
	}

	return nil
}
