package accountscontrollers

import (
	"os"

	"github.com/gofiber/fiber/v3"
	accountsusecases "github.com/henrique998/go-auth/internal/app/usecases/accounts-usecases"
	"github.com/henrique998/go-auth/internal/infra/database"
	"github.com/henrique998/go-auth/internal/infra/database/repositories"
	"github.com/henrique998/go-auth/internal/infra/providers"
	"github.com/henrique998/go-auth/internal/infra/utils"
	"github.com/twilio/twilio-go"
)

func Send2FACodeController(c fiber.Ctx) error {
	cookie := c.Cookies("goauth:access_token")

	if cookie == "" {
		return c.Status(fiber.StatusBadRequest).SendString("Cookie not found!")
	}

	accountId, err := utils.ParseJWTToken(cookie, os.Getenv("JWT_SECRET"))
	if err != nil {
		return c.Status(err.GetStatus()).SendString(err.GetMessage())
	}

	db := database.ConnectToDb()
	twilioClient := twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: os.Getenv("TWILIO_ACCOUNT_SID"),
		Password: os.Getenv("TWILIO_AUTH_TOKEN"),
	})
	defer db.Close()

	repo := repositories.PGAccountsRepository{Db: db}
	vtRepo := repositories.PGVerificationTokensRepository{Db: db}
	twoFactorAuthProvider := providers.TwilioTwoFactorAuthProvider{
		Client: twilioClient,
	}

	usecase := accountsusecases.Send2faCodeUseCase{
		Repo:                  &repo,
		VTRepo:                &vtRepo,
		TwoFactorAuthProvider: &twoFactorAuthProvider,
	}

	err = usecase.Execute(accountId)
	if err != nil {
		return c.Status(err.GetStatus()).SendString(err.GetMessage())
	}

	return c.SendString("Code sent succesfully!")
}
