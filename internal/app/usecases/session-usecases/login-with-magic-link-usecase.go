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

type LoginWithMagicLinkUseCase struct {
	Repo          contracts.AccountsRepository
	MLRepo        contracts.MagicLinksRepository
	DevicesRepo   contracts.DevicesRepository
	ATProvider    contracts.AuthTokensProvider
	EmailProvider contracts.EmailProvider
}

func (uc *LoginWithMagicLinkUseCase) Execute(req request.LoginWithMagicLinkRequest) (string, string, errors.IAppError) {
	magicLink := uc.MLRepo.FindByValue(req.Code)
	if magicLink == nil {
		return "", "", errors.NewAppError("Code not found!", http.StatusNotFound)
	}

	if time.Now().After(magicLink.ExpiresAt) {
		return "", "", errors.NewAppError("Code has expired!", http.StatusUnauthorized)
	}

	account := uc.Repo.FindById(magicLink.AccountId)
	if account == nil {
		return "", "", errors.NewAppError("Account not found!", http.StatusNotFound)
	}

	accessToken, refreshToken, tokenErr := uc.ATProvider.GenerateAuthTokens(account.ID)
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

	country, city, geoErr := utils.GetGeoLocation(req.IP)
	if geoErr != nil {
		logger.Error("Error trying to retrive geolocation data", geoErr)
		return "", "", errors.NewAppError("internal server error", 500)
	}

	if lastCountry != "" && (lastCountry != country || lastCity != city) {
		msg := "Sua conta foi acessada em outra localização. caso não tenha sido você recomendamos que altere sua senha. obrigado pela atenção!"
		uc.EmailProvider.SendMail(account.Email, "login suspeito.", msg)
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

		err := uc.DevicesRepo.Create(*device)
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

	err := uc.Repo.Update(*account)
	if err != nil {
		logger.Error("Error trying to update account data.", err)
		return "", "", errors.NewAppError("internal server error.", 500)
	}

	err = uc.MLRepo.Delete(magicLink.ID)
	if err != nil {
		logger.Error("Error while delete magic link", err)
		return "", "", errors.NewAppError("internal server error", http.StatusInternalServerError)
	}

	return accessToken, refreshToken, nil
}
