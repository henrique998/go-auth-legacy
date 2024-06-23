package contracts

import (
	"github.com/henrique998/go-auth/internal/app/entities"
	"github.com/henrique998/go-auth/internal/app/errors"
)

type AccountsRepository interface {
	FindById(accountId string) (*entities.Account, errors.IAppError)
	FindByEmail(email string) (*entities.Account, errors.IAppError)
	Create(a entities.Account) errors.IAppError
	Update(a entities.Account) errors.IAppError
}
