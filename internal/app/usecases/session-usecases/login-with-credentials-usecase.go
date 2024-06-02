package sessionusecases

import (
	"os"
	"time"

	"github.com/henrique998/go-auth/internal/app/contracts"
	"github.com/henrique998/go-auth/internal/app/errors"
	"github.com/henrique998/go-auth/internal/app/request"
	"github.com/henrique998/go-auth/internal/configs/logger"
	"github.com/henrique998/go-auth/internal/infra/utils"
)

type LoginWithCredentialsUseCase struct {
	Repo          contracts.AccountsRepository
	EmailProvider contracts.EmailProvider
}

func (uc *LoginWithCredentialsUseCase) Execute(req request.LoginWithCredentialsRequest) (string, errors.IAppError) {
	account, accountErr := uc.Repo.FindByEmail(req.Email)
	if account == nil || accountErr != nil {
		return "", errors.NewAppError("email or password incorrect!", 400)
	}

	passwordMatch := utils.ComparePassword(req.Pass, account.Pass)
	if !passwordMatch {
		return "", errors.NewAppError("email or password incorrect!", 400)
	}

	token, err := utils.GenerateJWTToken(account.ID, os.Getenv("JWT_SECRET"))
	if err != nil {
		logger.Error("Error trying to generate jwt token", err)
		return "", errors.NewAppError("internal server error.", 500)
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
		return "", errors.NewAppError("internal server error", 500)
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
		return "", errors.NewAppError("internal server error.", 500)
	}

	return token, nil
}
