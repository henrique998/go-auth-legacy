package contracts

import "github.com/henrique998/go-auth/internal/app/errors"

type AuthTokensProvider interface {
	GenerateAuthTokens(accountId string) (string, string, errors.IAppError)
	ValidateJWTToken(token string) (string, errors.IAppError)
}
