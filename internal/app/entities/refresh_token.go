package entities

import (
	"time"

	"github.com/google/uuid"
)

type RefreshToken struct {
	ID        string
	Value     string
	AccountId string
	ExpiresAt time.Time
	CreatedAt time.Time
}

func NewRefreshToken(val, accountId string, expiresAt time.Time) *RefreshToken {
	return &RefreshToken{
		ID:        uuid.New().String(),
		Value:     val,
		AccountId: accountId,
		ExpiresAt: expiresAt,
		CreatedAt: time.Now(),
	}
}

func NewExistingRefreshToken(id, val, accountId string, expiresAt, createdAt time.Time) *RefreshToken {
	return &RefreshToken{
		ID:        id,
		Value:     val,
		AccountId: accountId,
		ExpiresAt: expiresAt,
		CreatedAt: createdAt,
	}
}
