package entities

import (
	"time"

	"github.com/henrique998/go-auth/internal/infra/utils"
)

type Device struct {
	ID          string     `json:"id"`
	AccountID   string     `json:"account_id"`
	DeviceName  string     `json:"device_name"`
	UserAgent   string     `json:"user_agent"`
	Platform    string     `json:"platform"`
	IPAddress   string     `json:"ip_address"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at"`
	LastLoginAt time.Time  `json:"last_login_at"`
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
