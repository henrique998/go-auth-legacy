package contracts

import (
	"github.com/henrique998/go-auth/internal/app/entities"
)

type AccountsRepository interface {
	FindById(accountId string) (*entities.Account, error)
	FindByEmail(email string) (*entities.Account, error)
	Create(a entities.Account) error
	Update(a entities.Account) error
}
