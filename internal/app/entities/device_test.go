package entities

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestDevice_NewDevice(t *testing.T) {
	assert := assert.New(t)

	accountId := "account-id"
	deviceName := "iPhone"
	userAgent := "Mozilla/5.0"
	platform := "iOS"
	ip := "192.168.0.1"
	lastLoginAt := time.Now().Add(-1 * time.Hour)

	device := NewDevice(accountId, deviceName, userAgent, platform, ip, lastLoginAt)

	assert.NotNil(device)
	assert.Equal(accountId, device.AccountID)
	assert.Equal(deviceName, device.DeviceName)
	assert.Equal(userAgent, device.UserAgent)
	assert.Equal(platform, device.Platform)
	assert.Equal(ip, device.IPAddress)
	assert.Equal(lastLoginAt, device.LastLoginAt)
	assert.WithinDuration(time.Now(), device.CreatedAt, time.Second)
	assert.Nil(device.UpdatedAt)
}

func TestDevice_NewExistingDevice(t *testing.T) {
	assert := assert.New(t)

	id := "existing-id"
	accountId := "account-id"
	deviceName := "iPhone"
	userAgent := "Mozilla/5.0"
	platform := "iOS"
	ip := "192.168.0.1"
	createdAt := time.Now().Add(-24 * time.Hour)
	lastLoginAt := time.Now().Add(-1 * time.Hour)
	updatedAt := time.Now().Add(-30 * time.Minute)

	device := NewExistingDevice(id, accountId, deviceName, userAgent, platform, ip, createdAt, lastLoginAt, &updatedAt)

	assert.NotNil(device)
	assert.Equal(id, device.ID)
	assert.Equal(accountId, device.AccountID)
	assert.Equal(deviceName, device.DeviceName)
	assert.Equal(userAgent, device.UserAgent)
	assert.Equal(platform, device.Platform)
	assert.Equal(ip, device.IPAddress)
	assert.Equal(createdAt, device.CreatedAt)
	assert.Equal(lastLoginAt, device.LastLoginAt)
	assert.Equal(updatedAt, *device.UpdatedAt)
}
