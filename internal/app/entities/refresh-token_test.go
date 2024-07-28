package entities

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestRefreshToken_NewRefreshToken(t *testing.T) {
	assert := assert.New(t)

	val := "some-token-value"
	accountId := "account-id"
	expiresAt := time.Now().Add(24 * time.Hour)

	refreshToken := NewRefreshToken(val, accountId, expiresAt)

	assert.NotNil(refreshToken)
	assert.Equal(val, refreshToken.Value)
	assert.Equal(accountId, refreshToken.AccountId)
	assert.Equal(expiresAt, refreshToken.ExpiresAt)
	assert.WithinDuration(time.Now(), refreshToken.CreatedAt, time.Second)
}

func TestRefreshToken_NewExistingRefreshToken(t *testing.T) {
	assert := assert.New(t)

	id := "existing-id"
	val := "some-token-value"
	accountId := "account-id"
	expiresAt := time.Now().Add(24 * time.Hour)
	createdAt := time.Now().Add(-24 * time.Hour)

	refreshToken := NewExistingRefreshToken(id, val, accountId, expiresAt, createdAt)

	assert.NotNil(refreshToken)
	assert.Equal(id, refreshToken.ID)
	assert.Equal(val, refreshToken.Value)
	assert.Equal(accountId, refreshToken.AccountId)
	assert.Equal(expiresAt, refreshToken.ExpiresAt)
	assert.Equal(createdAt, refreshToken.CreatedAt)
}
