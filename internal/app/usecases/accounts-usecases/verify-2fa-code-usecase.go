package accountsusecases

import (
	"net/http"
	"time"

	"github.com/henrique998/go-auth/internal/app/contracts"
	appError "github.com/henrique998/go-auth/internal/app/errors"
	"github.com/henrique998/go-auth/internal/app/request"
	"github.com/henrique998/go-auth/internal/configs/logger"
)

type Verify2faCodeUseCase struct {
	Repo   contracts.AccountsRepository
	VTRepo contracts.VerificationCodesRepository
}

func (uc *Verify2faCodeUseCase) Execute(req request.Verify2faRequest) appError.IAppError {
	logger.Info("Init Verify2faCode UseCase")

	verificationCode := uc.VTRepo.FindByValue(req.Code)

	if verificationCode == nil {
		return appError.NewAppError("verification code not found", http.StatusNotFound)
	}

	if verificationCode.AccountId != req.AccountId {
		return appError.NewAppError("unauthorized action", http.StatusUnauthorized)
	}

	now := time.Now()

	if verificationCode.ExpiresAt.Before(now) {
		return appError.NewAppError("verification code has expired", http.StatusUnauthorized)
	}

	account := uc.Repo.FindById(req.AccountId)

	if account.Is2faEnabled {
		return appError.NewAppError("Two factor authentication already carried out", http.StatusUnauthorized)
	}

	account.Is2faEnabled = true
	account.UpdatedAt = &now

	uc.Repo.Update(*account)

	uc.VTRepo.Delete(verificationCode.ID)

	return nil
}
