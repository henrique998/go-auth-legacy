package accountsusecases

import (
	"net/http"
	"time"

	"github.com/henrique998/go-auth/internal/app/contracts"
	appError "github.com/henrique998/go-auth/internal/app/errors"
	"github.com/henrique998/go-auth/internal/app/request"
	"github.com/henrique998/go-auth/internal/configs/logger"
	"github.com/henrique998/go-auth/internal/infra/utils"
)

type UpdatePassUsecase struct {
	Repo   contracts.AccountsRepository
	VTRepo contracts.VerificationTokensRepository
}

func (uc *UpdatePassUsecase) Execute(req request.NewPassRequest) appError.IAppError {
	logger.Info("Init UpdatePass UseCase")

	code := uc.VTRepo.FindByValue(req.Code)

	if code == nil {
		return appError.NewAppError("code not found.", http.StatusNotFound)
	}

	account := uc.Repo.FindById(code.AccountId)

	now := time.Now()

	if now.After(code.ExpiresAt) {
		return appError.NewAppError("code has expired.", http.StatusUnauthorized)
	}

	if req.NewPass != req.NewPassConfirmation {
		return appError.NewAppError("new password and confirmation must be equals.", http.StatusBadRequest)
	}

	if len(req.NewPass) < 6 {
		return appError.NewAppError("new password must contain 6 or more characters.", http.StatusBadRequest)
	}

	if utils.ComparePassword(req.NewPass, *account.Pass) {
		return appError.NewAppError("new password cannot be the same as the previous one.", http.StatusBadRequest)
	}

	newPassHash, err := utils.HashPass(req.NewPass)
	if err != nil {
		logger.Error("Error trying to hash new password", err)
		return appError.NewAppError("internal server error.", http.StatusInternalServerError)
	}

	account.Pass = &newPassHash
	account.UpdatedAt = &now

	uc.Repo.Update(*account)
	uc.VTRepo.Delete(code.ID)

	return nil
}
