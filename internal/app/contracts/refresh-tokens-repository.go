package contracts

import (
	"github.com/henrique998/go-auth/internal/app/entities"
)

type RefreshTokensRepository interface {
	FindByValue(val string) (*entities.RefreshToken, error)
	Create(rt entities.RefreshToken) error
	Delete(val string) error
}
