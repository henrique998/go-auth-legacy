package contracts

import (
	"github.com/henrique998/go-setup/internal/app/entities"
	"github.com/henrique998/go-setup/internal/app/errors"
)

type UsersRepository interface {
	Create(u entities.IUser) errors.IAppError
	FindByEmail(email string) (*entities.IUser, errors.IAppError)
}
