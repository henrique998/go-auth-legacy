package entities

import (
	"time"

	"github.com/henrique998/go-auth/internal/infra/utils"
)

type LoginAttempt struct {
	ID          string    `json:"id"`
	Email       string    `json:"email"`
	IPAddress   string    `json:"ip_address"`
	UserAgent   string    `json:"user_agent"`
	Success     bool      `json:"success"`
	AttemptedAt time.Time `json:"attempted_at"`
}

func NewLoginAttempt(email, ip, userAgent string, success bool) *LoginAttempt {
	return &LoginAttempt{
		ID:          utils.GenerateUUID(),
		Email:       email,
		IPAddress:   ip,
		UserAgent:   userAgent,
		Success:     success,
		AttemptedAt: time.Now(),
	}
}
