package contracts

import (
	"github.com/henrique998/go-auth/internal/app/entities"
)

type VerificationTokensRepository interface {
	FindByValue(val string) *entities.VerificationToken
	Create(vt entities.VerificationToken) error
	Delete(tokenId string) error
}
