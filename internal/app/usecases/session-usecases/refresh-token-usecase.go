package sessionusecases

import (
	"github.com/henrique998/go-auth/internal/app/contracts"
	"github.com/henrique998/go-auth/internal/app/errors"
	"github.com/henrique998/go-auth/internal/configs/logger"
)

type RefreshTokenUseCase struct {
	Repo   contracts.RefreshTokensRepository
	ATRepo contracts.AuthTokensProvider
}

func (uc *RefreshTokenUseCase) Execute(refreshToken string) (string, string, errors.IAppError) {
	logger.Info("Init RefreshToken UseCase")

	accountId, err := uc.ATRepo.ValidateJWTToken(refreshToken)
	if err != nil {
		return "", "", err
	}

	newAccessToken, newRefreshToken, err := uc.ATRepo.GenerateAuthTokens(accountId)
	if err != nil {
		return "", "", err
	}

	uc.Repo.Delete(refreshToken)

	return newAccessToken, newRefreshToken, nil
}
