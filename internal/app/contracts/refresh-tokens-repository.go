package contracts

import (
	"github.com/henrique998/go-auth/internal/app/entities"
	"github.com/henrique998/go-auth/internal/app/errors"
)

type RefreshTokensRepository interface {
	FindByValue(val string) (*entities.RefreshToken, errors.IAppError)
	Create(rt entities.RefreshToken) errors.IAppError
}
