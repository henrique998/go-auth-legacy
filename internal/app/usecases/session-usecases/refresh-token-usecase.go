package sessionusecases

import (
	"github.com/henrique998/go-auth/internal/app/contracts"
	"github.com/henrique998/go-auth/internal/app/errors"
	"github.com/henrique998/go-auth/internal/app/response"
	"github.com/henrique998/go-auth/internal/configs/logger"
	"github.com/henrique998/go-auth/internal/infra/utils"
)

type RefreshTokenUseCase struct {
	Repo   contracts.RefreshTokensRepository
	ATRepo contracts.AuthTokensProvider
}

func (uc *RefreshTokenUseCase) Execute(refreshToken string) (response.AuthTokensResponse, errors.IAppError) {
	logger.Info("Init RefreshToken UseCase")

	accountId, err := utils.ValidateJWTToken(refreshToken)
	if err != nil {
		return response.AuthTokensResponse{}, err
	}

	newAccessToken, newRefreshToken, err := uc.ATRepo.GenerateAuthTokens(accountId)
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
