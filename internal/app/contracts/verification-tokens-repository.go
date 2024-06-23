package contracts

import (
	"github.com/henrique998/go-auth/internal/app/entities"
)

type VerificationTokensRepository interface {
	FindByValue(val string) (*entities.VerificationToken, error)
	Create(vt entities.VerificationToken) error
	Delete(tokenId string) error
}
