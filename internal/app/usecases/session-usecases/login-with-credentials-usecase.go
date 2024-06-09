package sessionusecases

import (
	"os"
	"time"

	"github.com/henrique998/go-auth/internal/app/contracts"
	"github.com/henrique998/go-auth/internal/app/entities"
	"github.com/henrique998/go-auth/internal/app/errors"
	"github.com/henrique998/go-auth/internal/app/request"
	"github.com/henrique998/go-auth/internal/configs/logger"
	"github.com/henrique998/go-auth/internal/infra/utils"
)

type LoginWithCredentialsUseCase struct {
	Repo          contracts.AccountsRepository
	RTRepo        contracts.RefreshTokensRepository
	EmailProvider contracts.EmailProvider
}

func (uc *LoginWithCredentialsUseCase) Execute(req request.LoginWithCredentialsRequest) (string, string, errors.IAppError) {
	account, accountErr := uc.Repo.FindByEmail(req.Email)
	if account == nil || accountErr != nil {
		return "", "", errors.NewAppError("email or password incorrect!", 400)
	}

	passwordMatch := utils.ComparePassword(req.Pass, account.Pass)
	if !passwordMatch {
		return "", "", errors.NewAppError("email or password incorrect!", 400)
	}

	accessToken, refreshToken, tokenErr := uc.generateAuthTokens(account.ID)
	if tokenErr != nil {
		return "", "", tokenErr
	}

	var lastCountry, lastCity string
	if account.LastLoginCountry != nil {
		lastCountry = *account.LastLoginCountry
	}
	if account.LastLoginCity != nil {
		lastCity = *account.LastLoginCity
	}

	country, city, err := utils.GetGeoLocation(req.IP)
	if err != nil {
		logger.Error("Error trying to retrive geolocation data", err)
		return "", "", errors.NewAppError("internal server error", 500)
	}

	if lastCountry != "" && (lastCountry != country || lastCity != city) {
		msg := "Sua conta foi acessada em outra localização. caso não tenha sido você recomendamos que altere sua senha. obrigado pela atenção!"
		uc.EmailProvider.SendMail(req.Email, "login suspeito.", msg)
	}

	now := time.Now()
	account.LastLoginAt = &now
	account.LastLoginCountry = &country
	account.LastLoginCity = &city
	account.LastLoginIp = &req.IP

	accountErr = uc.Repo.Update(account)
	if accountErr != nil {
		logger.Error("Error trying to update account data.", err)
		return "", "", errors.NewAppError("internal server error.", 500)
	}

	return accessToken, refreshToken, nil
}

func (uc *LoginWithCredentialsUseCase) generateAuthTokens(accountId string) (string, string, errors.IAppError) {
	tokenExpiresAt := time.Now().Add(15 * time.Minute)
	accessToken, tokenErr := utils.GenerateJWTToken(accountId, tokenExpiresAt, os.Getenv("JWT_SECRET"))
	if tokenErr != nil {
		logger.Error("Error trying to generate access token token", tokenErr)
		return "", "", errors.NewAppError("internal server error.", 500)
	}

	refreshTokenExpiresAt := time.Now().Add(15 * time.Minute)
	refreshToken, tokenErr := utils.GenerateJWTToken(accountId, refreshTokenExpiresAt, os.Getenv("JWT_SECRET"))
	if tokenErr != nil {
		logger.Error("Error trying to generate refresh token", tokenErr)
		return "", "", errors.NewAppError("internal server error.", 500)
	}

	rt := entities.NewRefreshToken(refreshToken, accountId, refreshTokenExpiresAt)

	uc.RTRepo.Create(*rt)

	return accessToken, refreshToken, nil
}
