package accountsusecases

import (
	"time"

	"github.com/henrique998/go-auth/internal/app/contracts"
	appErrors "github.com/henrique998/go-auth/internal/app/errors"
	"github.com/henrique998/go-auth/internal/configs/logger"
)

type VerifyEmailUseCase struct {
	Repo   contracts.AccountsRepository
	VTRepo contracts.VerificationTokensRepository
}

func (uc *VerifyEmailUseCase) Execute(token string) appErrors.IAppError {
	logger.Info("Init VerifyEmail UseCase")

	verificationToken, err := uc.VTRepo.FindByValue(token)
	if err != nil {
		return err
	}

	account, err := uc.Repo.FindById(verificationToken.AccountId)
	if err != nil {
		return err
	}

	if time.Now().After(verificationToken.ExpiresAt) {
		return appErrors.NewAppError("verification token has expired!", 400)
	}

	account.IsEmailVerified = true
	now := time.Now()
	account.UpdatedAt = &now

	uc.Repo.Update(account)

	uc.VTRepo.Delete(verificationToken.ID)

	return nil
}
