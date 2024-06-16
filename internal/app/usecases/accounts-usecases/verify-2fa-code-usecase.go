package accountsusecases

import (
	"errors"
	"net/http"
	"time"

	"github.com/henrique998/go-auth/internal/app/contracts"
	appError "github.com/henrique998/go-auth/internal/app/errors"
	"github.com/henrique998/go-auth/internal/app/request"
	"github.com/henrique998/go-auth/internal/configs/logger"
)

type Verify2faCodeUseCase struct {
	Repo   contracts.AccountsRepository
	VTRepo contracts.VerificationTokensRepository
}

func (uc *Verify2faCodeUseCase) Execute(req request.Verify2faRequest) appError.IAppError {
	logger.Info("Init Verify2faCode UseCase")

	verificationCode, err := uc.VTRepo.FindByValue(req.Code)
	if err != nil {
		logger.Error("Error while find verification code", errors.New(err.GetMessage()))
		return err
	}

	if verificationCode == nil {
		return appError.NewAppError("Verification code not found!", http.StatusNotFound)
	}

	if verificationCode.AccountId != req.AccountId {
		return appError.NewAppError("Unauthorized action!", http.StatusUnauthorized)
	}

	now := time.Now()

	if verificationCode.ExpiresAt.Before(now) {
		return appError.NewAppError("Verification code has expired", http.StatusUnauthorized)
	}

	account, err := uc.Repo.FindById(req.AccountId)
	if err != nil {
		return err
	}

	if account.Is2faEnabled {
		return appError.NewAppError("Two factor authentication already carried out!", http.StatusBadRequest)
	}

	account.Is2faEnabled = true
	account.UpdatedAt = &now

	uc.Repo.Update(account)

	uc.VTRepo.Delete(verificationCode.ID)

	return nil
}
