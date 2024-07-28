package entities

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestVerificationToken_NewVerificationToken(t *testing.T) {
	assert := assert.New(t)

	val := "verification-token-value"
	accountId := "account-id"
	expiresAt := time.Now().Add(24 * time.Hour)

	verificationToken := NewVerificationToken(val, accountId, expiresAt)

	assert.NotNil(verificationToken)
	assert.Equal(val, verificationToken.Value)
	assert.Equal(accountId, verificationToken.AccountId)
	assert.Equal(expiresAt, verificationToken.ExpiresAt)
	assert.WithinDuration(time.Now(), verificationToken.CreatedAt, time.Second)
}

func TestVerificationToken_NewExistingVerificationToken(t *testing.T) {
	assert := assert.New(t)

	id := "existing-id"
	val := "verification-token-value"
	accountId := "account-id"
	createdAt := time.Now().Add(-24 * time.Hour)
	expiresAt := time.Now().Add(24 * time.Hour)

	verificationToken := NewExistingerificationToken(id, val, accountId, createdAt, expiresAt)

	assert.NotNil(verificationToken)
	assert.Equal(id, verificationToken.ID)
	assert.Equal(val, verificationToken.Value)
	assert.Equal(accountId, verificationToken.AccountId)
	assert.Equal(createdAt, verificationToken.CreatedAt)
	assert.Equal(expiresAt, verificationToken.ExpiresAt)
}
