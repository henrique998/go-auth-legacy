package utils

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/henrique998/go-auth/internal/app/errors"
)

func ValidateJWTToken(tokenString string) (string, errors.IAppError) {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err != nil {
		return "", errors.NewAppError("invalid token", 401)
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return "", errors.NewAppError("invalid token", 401)
	}

	accountId := claims["sub"].(string)
	exp := int64(claims["exp"].(float64))

	if exp < time.Now().Unix() {
		return "", errors.NewAppError("token expired", 401)
	}

	return accountId, nil
}
