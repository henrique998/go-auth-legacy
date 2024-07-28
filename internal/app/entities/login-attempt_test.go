package entities

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestLoginAttempt_NewLoginAttempt(t *testing.T) {
	assert := assert.New(t)

	email := "test@example.com"
	ip := "192.168.0.1"
	userAgent := "Mozilla/5.0"
	success := true

	loginAttempt := NewLoginAttempt(email, ip, userAgent, success)

	assert.NotNil(loginAttempt)
	assert.Equal(email, loginAttempt.Email)
	assert.Equal(ip, loginAttempt.IPAddress)
	assert.Equal(userAgent, loginAttempt.UserAgent)
	assert.Equal(success, loginAttempt.Success)
	assert.WithinDuration(time.Now(), loginAttempt.AttemptedAt, time.Second)
}
