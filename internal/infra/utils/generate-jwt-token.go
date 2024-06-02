package utils

import "github.com/golang-jwt/jwt/v5"

func GenerateJWTToken(sub string, secret string) (string, error) {
	claims := jwt.MapClaims{
		"sub": sub,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(secret))
}
