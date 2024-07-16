package accountsusecases

import (
	"net/http"

	"github.com/henrique998/go-auth/internal/app/contracts"
	"github.com/henrique998/go-auth/internal/app/entities"
	appError "github.com/henrique998/go-auth/internal/app/errors"
)

type GetAccountDevicesUseCase struct {
	Repo        contracts.AccountsRepository
	DevicesRepo contracts.DevicesRepository
}

func (uc *GetAccountDevicesUseCase) Execute(accountId string) ([]entities.Device, appError.IAppError) {
	account := uc.Repo.FindById(accountId)

	if account == nil {
		return nil, appError.NewAppError("account does not exists", http.StatusNotFound)
	}

	devices := uc.DevicesRepo.FindManyByAccountId(accountId)

	return devices, nil
}
