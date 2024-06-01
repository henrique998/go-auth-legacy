package entities

import (
	"time"

	"github.com/google/uuid"
)

type VerificationToken struct {
	ID        string
	Value     string
	AccountId string
	CreatedAt time.Time
	ExpiresAt time.Time
}

func NewVerificationToken(val, accountId string, expiresAt time.Time) *VerificationToken {
	return &VerificationToken{
		ID:        uuid.New().String(),
		Value:     val,
		AccountId: accountId,
		CreatedAt: time.Now(),
		ExpiresAt: expiresAt,
	}
}

func NewExistingerificationToken(id, val, accountId string, createdAt, expiresAt time.Time) *VerificationToken {
	return &VerificationToken{
		ID:        id,
		Value:     val,
		AccountId: accountId,
		CreatedAt: createdAt,
		ExpiresAt: expiresAt,
	}
}
