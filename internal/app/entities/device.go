package entities

import (
	"time"

	"github.com/henrique998/go-auth/internal/infra/utils"
)

type Device struct {
	ID          string
	AccountID   string
	DeviceName  string
	UserAgent   string
	Platform    string
	IPAddress   string
	CreatedAt   time.Time
	UpdatedAt   *time.Time
	LastLoginAt time.Time
}

func NewDevice(accountId, deviceName, userAgent, platform, ip string, lastLoginAt time.Time) *Device {
	return &Device{
		ID:          utils.GenerateUUID(),
		AccountID:   accountId,
		DeviceName:  deviceName,
		UserAgent:   userAgent,
		Platform:    platform,
		IPAddress:   ip,
		CreatedAt:   time.Now(),
		UpdatedAt:   nil,
		LastLoginAt: lastLoginAt,
	}
}

func NewExistingDevice(id, accountId, deviceName, userAgent, platform, ip string, createdAt, lastLoginAt time.Time, updatedtAt *time.Time) *Device {
	return &Device{
		ID:          id,
		AccountID:   accountId,
		DeviceName:  deviceName,
		UserAgent:   userAgent,
		Platform:    platform,
		IPAddress:   ip,
		CreatedAt:   createdAt,
		UpdatedAt:   updatedtAt,
		LastLoginAt: lastLoginAt,
	}
}
