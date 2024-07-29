package entities

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestVerificationCode_NewVerificationCode(t *testing.T) {
	assert := assert.New(t)

	val := "verification-code-value"
	accountId := "account-id"
	expiresAt := time.Now().Add(24 * time.Hour)

	verificationCode := NewVerificationCode(val, accountId, expiresAt)

	assert.NotNil(verificationCode)
	assert.Equal(val, verificationCode.Value)
	assert.Equal(accountId, verificationCode.AccountId)
	assert.Equal(expiresAt, verificationCode.ExpiresAt)
	assert.WithinDuration(time.Now(), verificationCode.CreatedAt, time.Second)
}

func TestVerificationCode_NewExistingVerificationCode(t *testing.T) {
	assert := assert.New(t)

	id := "existing-id"
	val := "verification-code-value"
	accountId := "account-id"
	createdAt := time.Now().Add(-24 * time.Hour)
	expiresAt := time.Now().Add(24 * time.Hour)

	verificationCode := NewExistingerificationCode(id, val, accountId, createdAt, expiresAt)

	assert.NotNil(verificationCode)
	assert.Equal(id, verificationCode.ID)
	assert.Equal(val, verificationCode.Value)
	assert.Equal(accountId, verificationCode.AccountId)
	assert.Equal(createdAt, verificationCode.CreatedAt)
	assert.Equal(expiresAt, verificationCode.ExpiresAt)
}
