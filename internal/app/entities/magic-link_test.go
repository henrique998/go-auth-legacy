package entities

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestMagicLink_NewMagicLink(t *testing.T) {
	assert := assert.New(t)

	accountId := "account-id"
	code := "magic-code"
	expiresAt := time.Now().Add(24 * time.Hour)

	magicLink := NewMagicLink(accountId, code, expiresAt)

	assert.NotNil(magicLink)
	assert.Equal(accountId, magicLink.AccountId)
	assert.Equal(code, magicLink.Code)
	assert.Equal(expiresAt, magicLink.ExpiresAt)
	assert.WithinDuration(time.Now(), magicLink.CreatedAt, time.Second)
}

func TestMagicLink_NewExistingMagicLink(t *testing.T) {
	assert := assert.New(t)

	id := "existing-id"
	accountId := "account-id"
	code := "magic-code"
	expiresAt := time.Now().Add(24 * time.Hour)
	createdAt := time.Now().Add(-24 * time.Hour)

	magicLink := NewExistingMagicLink(id, accountId, code, expiresAt, createdAt)

	assert.NotNil(magicLink)
	assert.Equal(id, magicLink.ID)
	assert.Equal(accountId, magicLink.AccountId)
	assert.Equal(code, magicLink.Code)
	assert.Equal(expiresAt, magicLink.ExpiresAt)
	assert.Equal(createdAt, magicLink.CreatedAt)
}
