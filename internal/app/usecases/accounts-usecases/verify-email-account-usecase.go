package accountsusecases

import (
	"net/http"
	"time"

	"github.com/henrique998/go-auth/internal/app/contracts"
	appError "github.com/henrique998/go-auth/internal/app/errors"
	"github.com/henrique998/go-auth/internal/configs/logger"
)

type VerifyEmailUseCase struct {
	Repo   contracts.AccountsRepository
	VTRepo contracts.VerificationCodesRepository
}

func (uc *VerifyEmailUseCase) Execute(token string) appError.IAppError {
	logger.Info("Init VerifyEmail UseCase")

	verificationToken := uc.VTRepo.FindByValue(token)

	if time.Now().After(verificationToken.ExpiresAt) {
		return appError.NewAppError("verification code has expired", http.StatusUnauthorized)
	}

	account := uc.Repo.FindById(verificationToken.AccountId)

	if account.IsEmailVerified {
		return appError.NewAppError("the email has already been verified", http.StatusUnauthorized)
	}

	account.IsEmailVerified = true
	now := time.Now()
	account.UpdatedAt = &now

	uc.Repo.Update(*account)

	uc.VTRepo.Delete(verificationToken.ID)

	return nil
}
