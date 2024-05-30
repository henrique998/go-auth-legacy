package contracts

import (
	"github.com/henrique998/go-auth/internal/app/entities"
	"github.com/henrique998/go-auth/internal/app/errors"
)

type AccountsRepository interface {
	Create(u entities.IAccount) errors.IAppError
	FindByEmail(email string) *entities.IAccount
}
