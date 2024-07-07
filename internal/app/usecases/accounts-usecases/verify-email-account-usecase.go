package accountsusecases

import (
	"time"

	"github.com/henrique998/go-auth/internal/app/contracts"
	appError "github.com/henrique998/go-auth/internal/app/errors"
	"github.com/henrique998/go-auth/internal/configs/logger"
)

type VerifyEmailUseCase struct {
	Repo   contracts.AccountsRepository
	VTRepo contracts.VerificationTokensRepository
}

func (uc *VerifyEmailUseCase) Execute(token string) appError.IAppError {
	logger.Info("Init VerifyEmail UseCase")

	verificationToken := uc.VTRepo.FindByValue(token)

	account := uc.Repo.FindById(verificationToken.AccountId)

	if time.Now().After(verificationToken.ExpiresAt) {
		return appError.NewAppError("verification token has expired!", 400)
	}

	account.IsEmailVerified = true
	now := time.Now()
	account.UpdatedAt = &now

	uc.Repo.Update(*account)

	uc.VTRepo.Delete(verificationToken.ID)

	return nil
}
