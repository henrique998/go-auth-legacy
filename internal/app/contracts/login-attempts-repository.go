package contracts

import "github.com/henrique998/go-auth/internal/app/entities"

type LoginAttemptsRepository interface {
	Create(loginAttempt entities.LoginAttempt) error
}
