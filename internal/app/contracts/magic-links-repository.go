package contracts

import (
	"github.com/henrique998/go-auth/internal/app/entities"
)

type MagicLinksRepository interface {
	FindByValue(val string) *entities.MagicLink
	Create(ml entities.MagicLink) error
	Delete(mlId string) error
}
