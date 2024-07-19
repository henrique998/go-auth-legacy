package sessionusecases

import (
	"net/http"
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
	DevicesRepo   contracts.DevicesRepository
	LARepository  contracts.LoginAttemptsRepository
	EmailProvider contracts.EmailProvider
	AtProvider    contracts.AuthTokensProvider
}

func (uc *LoginWithCredentialsUseCase) Execute(req request.LoginWithCredentialsRequest) (string, string, errors.IAppError) {
	logger.Info("Init LoginWithCredentials UseCase")

	account := uc.Repo.FindByEmail(req.Email)
	if account == nil {
		la := entities.NewLoginAttempt(req.Email, req.IP, req.UserAgent, false)

		err := uc.LARepository.Create(*la)
		if err != nil {
			logger.Error("Error trying to create login attempt record.", err)
		}

		return "", "", errors.NewAppError("email or password incorrect!", 400)
	}

	if account.Pass == nil {
		la := entities.NewLoginAttempt(req.Email, req.IP, req.UserAgent, false)

		err := uc.LARepository.Create(*la)
		if err != nil {
			logger.Error("Error trying to create login attempt record.", err)
		}

		return "", "", errors.NewAppError("Login method not allowed!", http.StatusUnauthorized)
	}

	passwordMatch := utils.ComparePassword(req.Pass, *account.Pass)
	if !passwordMatch {
		la := entities.NewLoginAttempt(req.Email, req.IP, req.UserAgent, false)

		err := uc.LARepository.Create(*la)
		if err != nil {
			logger.Error("Error trying to create login attempt record.", err)
		}

		return "", "", errors.NewAppError("email or password incorrect!", 400)
	}

	if !account.IsEmailVerified {
		la := entities.NewLoginAttempt(req.Email, req.IP, req.UserAgent, false)

		err := uc.LARepository.Create(*la)
		if err != nil {
			logger.Error("Error trying to create login attempt record.", err)
		}

		return "", "", errors.NewAppError("Only verified accounts can log in!", http.StatusUnauthorized)
	}

	accessToken, refreshToken, tokenErr := uc.AtProvider.GenerateAuthTokens(account.ID)
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

	deviceDetails := utils.GetDeviceDetails(req.UserAgent)
	now := time.Now()

	device := uc.DevicesRepo.FindByIpAndAccountId(req.IP, account.ID)

	if device == nil {
		device = entities.NewDevice(
			account.ID,
			deviceDetails.Name,
			req.UserAgent,
			deviceDetails.Platform,
			req.IP,
			now,
		)

		err = uc.DevicesRepo.Create(*device)
		if err != nil {
			logger.Error("Error trying to create device", err)
			return "", "", errors.NewAppError("internal server error", 500)
		}
	} else {
		device.LastLoginAt = now
		uc.DevicesRepo.Update(*device)
	}

	account.LastLoginAt = &now
	account.LastLoginCountry = &country
	account.LastLoginCity = &city
	account.LastLoginIp = &req.IP

	err = uc.Repo.Update(*account)
	if err != nil {
		logger.Error("Error trying to update account data.", err)
		return "", "", errors.NewAppError("internal server error.", 500)
	}

	la := entities.NewLoginAttempt(account.Email, req.IP, req.UserAgent, true)

	err = uc.LARepository.Create(*la)
	if err != nil {
		logger.Error("Error trying to create login attempt record.", err)
		return accessToken, refreshToken, nil
	}

	return accessToken, refreshToken, nil
}
