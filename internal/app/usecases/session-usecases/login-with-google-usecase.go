package sessionusecases

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/henrique998/go-auth/internal/app/contracts"
	"github.com/henrique998/go-auth/internal/app/entities"
	"github.com/henrique998/go-auth/internal/app/errors"
	"github.com/henrique998/go-auth/internal/app/request"
	"github.com/henrique998/go-auth/internal/configs/logger"
	"github.com/henrique998/go-auth/internal/infra/providers"
	"github.com/henrique998/go-auth/internal/infra/utils"
)

type LoginWithGoogleUseCase struct {
	Repo          contracts.AccountsRepository
	EmailProvider contracts.EmailProvider
	AtProvider    contracts.AuthTokensProvider
	DevicesRepo   contracts.DevicesRepository
}

var userInfoUrl = "https://www.googleapis.com/oauth2/v2/userinfo?access_token="
var personFieldsUrl = "https://people.googleapis.com/v1/people/me?personFields=phoneNumbers"

func (uc *LoginWithGoogleUseCase) Execute(data request.LoginWithGoogleRequest) (string, string, errors.IAppError) {
	logger.Info("Init LoginWithGoogle UseCase")

	accessToken, err := utils.GetGoogleAccessToken(data.Code)
	if err != nil {
		logger.Error("Error trying while get google oauth access token", err)
		return "", "", errors.NewAppError("internal server error", 500)
	}

	res, err := http.Get(userInfoUrl + accessToken)
	if err != nil {
		logger.Error("Error trying while get user info using access token", err)
		return "", "", errors.NewAppError("internal server error", 500)
	}
	defer res.Body.Close()

	userInfo := make(map[string]any)
	if err := json.NewDecoder(res.Body).Decode(&userInfo); err != nil {
		logger.Error("failed to decode user info:", err)
		return "", "", errors.NewAppError("internal server error", 500)
	}

	req, err := http.NewRequest("GET", personFieldsUrl, nil)
	if err != nil {
		logger.Error("failed to generate http request:", err)
		return "", "", errors.NewAppError("internal server error", 500)
	}

	gcProvider := providers.GoogleHTTPClientProvider{
		Token: accessToken,
	}

	res, err = gcProvider.Do(req)
	if err != nil {
		logger.Error("failed while make http request:", err)
		return "", "", errors.NewAppError("internal server error", 500)
	}

	defer res.Body.Close()
	if err := json.NewDecoder(res.Body).Decode(&userInfo); err != nil {
		logger.Error("failed to decode user info:", err)
		return "", "", errors.NewAppError("internal server error", 500)
	}

	id := userInfo["id"].(string)
	name := userInfo["name"].(string)
	email := userInfo["email"].(string)

	account := uc.Repo.FindByEmail(email)

	if account == nil {
		account = entities.NewAccount(name, email, "", "", id)

		uc.Repo.Create(*account)
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

	country, city, err := utils.GetGeoLocation(data.IP)
	if err != nil {
		logger.Error("Error trying to retrive geolocation data", err)
		return "", "", errors.NewAppError("internal server error", 500)
	}

	if lastCountry != "" && (lastCountry != country || lastCity != city) {
		msg := "Sua conta foi acessada em outra localização. caso não tenha sido você recomendamos que altere sua senha. obrigado pela atenção!"
		uc.EmailProvider.SendMail(account.Email, "login suspeito.", msg)
	}

	deviceDetails := utils.GetDeviceDetails(data.UserAgent)
	now := time.Now()

	device := uc.DevicesRepo.FindByIpAndAccountId(data.IP, account.ID)

	if device == nil {
		device = entities.NewDevice(
			account.ID,
			deviceDetails.Name,
			data.UserAgent,
			deviceDetails.Platform,
			data.IP,
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
	account.LastLoginIp = &data.IP

	err = uc.Repo.Update(*account)
	if err != nil {
		logger.Error("Error trying to update account data.", err)
		return "", "", errors.NewAppError("internal server error.", 500)
	}

	return accessToken, refreshToken, nil
}
