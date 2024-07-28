package entities

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestAccount_NewAccount(t *testing.T) {
	assert := assert.New(t)

	name := "jhon"
	email := "jhondoe@gmail.com"
	pass := "123456"
	phone := "999999999"

	account := NewAccount(name, email, pass, phone, "")

	assert.NotNil(account)
	assert.Equal(name, account.Name)
	assert.Equal(email, account.Email)
	assert.Equal(pass, *account.Pass)
	assert.Equal(phone, *account.Phone)
}

func TestAccount_NewExistingAccount(t *testing.T) {
	assert := assert.New(t)

	id := "existing-id"
	name := "jhon"
	email := "jhondoe@gmail.com"
	pass := "123456"
	phone := "999999999"
	now := time.Now()
	lastLoginAt := now.Add(-1 * time.Hour)
	lastLoginIp := "0.0.0.0"
	lastLoginCountry := "br"
	lastLoginCity := "sp"

	account := NewExistingAccount(
		id,
		name,
		email,
		pass,
		phone,
		"",
		true,
		false,
		lastLoginAt,
		lastLoginIp,
		lastLoginCountry,
		lastLoginCity,
		now,
		now,
	)

	assert.NotNil(account)
	assert.Equal(id, account.ID)
	assert.Equal(name, account.Name)
	assert.Equal(email, account.Email)
	assert.Equal(pass, *account.Pass)
	assert.Equal(phone, *account.Phone)
	assert.True(account.Is2faEnabled)
	assert.False(account.IsEmailVerified)
}
