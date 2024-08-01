package providers

import (
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/henrique998/go-auth/internal/app/contracts"
	"github.com/henrique998/go-auth/internal/app/entities"
	"github.com/henrique998/go-auth/internal/app/errors"
	"github.com/henrique998/go-auth/internal/configs/logger"
	"github.com/henrique998/go-auth/internal/infra/utils"
)

type AuthTokensProvider struct {
	RTRepo contracts.RefreshTokensRepository
}

func (p *AuthTokensProvider) GenerateAuthTokens(accountId string) (string, string, errors.IAppError) {
	tokenExpiresAt := time.Now().Add(15 * time.Minute)
	accessToken, tokenErr := utils.GenerateJWTToken(accountId, tokenExpiresAt, os.Getenv("JWT_SECRET"))
	if tokenErr != nil {
		logger.Error("Error trying to generate access token token", tokenErr)
		return "", "", errors.NewAppError("internal server error.", http.StatusInternalServerError)
	}

	refreshTokenExpiresAt := time.Now().Add(time.Hour * 24 * 30)
	refreshToken, tokenErr := utils.GenerateJWTToken(accountId, refreshTokenExpiresAt, os.Getenv("JWT_SECRET"))
	if tokenErr != nil {
		logger.Error("Error trying to generate refresh token", tokenErr)
		return "", "", errors.NewAppError("internal server error.", http.StatusInternalServerError)
	}

	rt := entities.NewRefreshToken(refreshToken, accountId, refreshTokenExpiresAt)

	p.RTRepo.Create(*rt)

	return accessToken, refreshToken, nil
}

func (p *AuthTokensProvider) ValidateJWTToken(tokenString string) (string, errors.IAppError) {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err != nil {
		return "", errors.NewAppError("invalid token", http.StatusUnauthorized)
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return "", errors.NewAppError("invalid token", http.StatusUnauthorized)
	}

	accountId := claims["sub"].(string)
	exp := int64(claims["exp"].(float64))

	if exp < time.Now().Unix() {
		return "", errors.NewAppError("token expired", http.StatusUnauthorized)
	}

	return accountId, nil
}
