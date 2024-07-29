package entities

import (
	"time"

	"github.com/google/uuid"
)

type VerificationCode struct {
	ID        string
	Value     string
	AccountId string
	CreatedAt time.Time
	ExpiresAt time.Time
}

func NewVerificationCode(val, accountId string, expiresAt time.Time) *VerificationCode {
	return &VerificationCode{
		ID:        uuid.New().String(),
		Value:     val,
		AccountId: accountId,
		CreatedAt: time.Now(),
		ExpiresAt: expiresAt,
	}
}

func NewExistingerificationCode(id, val, accountId string, createdAt, expiresAt time.Time) *VerificationCode {
	return &VerificationCode{
		ID:        id,
		Value:     val,
		AccountId: accountId,
		CreatedAt: createdAt,
		ExpiresAt: expiresAt,
	}
}
