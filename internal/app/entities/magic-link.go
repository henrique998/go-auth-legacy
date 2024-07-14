package entities

import (
	"time"

	"github.com/henrique998/go-auth/internal/infra/utils"
)

type MagicLink struct {
	ID        string
	AccountId string
	Code      string
	ExpiresAt time.Time
	CreatedAt time.Time
}

func NewMagicLink(accountId, code string, expiresAt time.Time) *MagicLink {
	return &MagicLink{
		ID:        utils.GenerateUUID(),
		AccountId: accountId,
		Code:      code,
		ExpiresAt: expiresAt,
		CreatedAt: time.Now(),
	}
}

func NewExistingMagicLink(id, accountId, code string, expiresAt, createdAt time.Time) *MagicLink {
	return &MagicLink{
		ID:        id,
		AccountId: accountId,
		Code:      code,
		ExpiresAt: expiresAt,
		CreatedAt: createdAt,
	}
}
