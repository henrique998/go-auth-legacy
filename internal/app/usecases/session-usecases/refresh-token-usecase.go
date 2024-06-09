package sessionusecases

import (
	"os"
	"time"

	"github.com/henrique998/go-auth/internal/app/contracts"
	"github.com/henrique998/go-auth/internal/app/entities"
	"github.com/henrique998/go-auth/internal/app/errors"
	"github.com/henrique998/go-auth/internal/app/response"
	"github.com/henrique998/go-auth/internal/configs/logger"
	"github.com/henrique998/go-auth/internal/infra/utils"
)

type RefreshTokenUseCase struct {
	Repo contracts.RefreshTokensRepository
}

func (uc *RefreshTokenUseCase) Execute(refreshToken string) (response.AuthTokensResponse, errors.IAppError) {
	accountId, err := utils.ValidateJWTToken(refreshToken)
	if err != nil {
		return response.AuthTokensResponse{}, err
	}

	newAccessToken, newRefreshToken, err := uc.generateNewAuthTokens(accountId)
	if err != nil {
		return response.AuthTokensResponse{}, err
	}

	uc.Repo.Delete(refreshToken)

	res := response.AuthTokensResponse{
		AccessToken:  newAccessToken,
		RefreshToken: newRefreshToken,
	}

	return res, nil
}

func (uc *RefreshTokenUseCase) generateNewAuthTokens(accountId string) (string, string, errors.IAppError) {
	tokenExpiresAt := time.Now().Add(15 * time.Minute)
	accessToken, tokenErr := utils.GenerateJWTToken(accountId, tokenExpiresAt, os.Getenv("JWT_SECRET"))
	if tokenErr != nil {
		logger.Error("Error trying to generate access token token", tokenErr)
		return "", "", errors.NewAppError("internal server error.", 500)
	}

	refreshTokenExpiresAt := time.Now().Add(time.Hour * 24 * 30)
	refreshToken, tokenErr := utils.GenerateJWTToken(accountId, refreshTokenExpiresAt, os.Getenv("JWT_SECRET"))
	if tokenErr != nil {
		logger.Error("Error trying to generate refresh token", tokenErr)
		return "", "", errors.NewAppError("internal server error.", 500)
	}

	rt := entities.NewRefreshToken(refreshToken, accountId, refreshTokenExpiresAt)

	uc.Repo.Create(*rt)

	return accessToken, refreshToken, nil
}
