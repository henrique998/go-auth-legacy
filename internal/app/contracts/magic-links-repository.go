package contracts

import (
	"github.com/henrique998/go-auth/internal/app/entities"
)

type MagicLinksRepository interface {
	Create(ml entities.MagicLink) error
}
