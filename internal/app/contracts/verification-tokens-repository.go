package contracts

import (
	"github.com/henrique998/go-auth/internal/app/entities"
	"github.com/henrique998/go-auth/internal/app/errors"
)

type VerificationTokensRepository interface {
	FindByValue(val string) (*entities.VerificationToken, errors.IAppError)
	Create(vt entities.VerificationToken) errors.IAppError
	Delete(tokenId string) errors.IAppError
}
