package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	Sub string `json:"sub"`
	jwt.RegisteredClaims
}

func GenerateJWTToken(sub string, expiresAt time.Time, secret string) (string, error) {
	claims := Claims{
		Sub: sub,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expiresAt),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(secret))
}
