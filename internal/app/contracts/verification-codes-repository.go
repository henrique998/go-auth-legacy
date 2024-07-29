package contracts

import (
	"github.com/henrique998/go-auth/internal/app/entities"
)

type VerificationCodesRepository interface {
	FindByValue(val string) *entities.VerificationCode
	Create(vt entities.VerificationCode) error
	Delete(tokenId string) error
}
