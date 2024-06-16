package contracts

import "github.com/henrique998/go-auth/internal/app/errors"

type TwoFactorAuthProvider interface {
	Send(from, to, message string) errors.IAppError
}
