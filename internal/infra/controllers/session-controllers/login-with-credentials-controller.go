package sessioncontrollers

import (
	"encoding/json"
	"os"

	"github.com/gofiber/fiber/v3"
	"github.com/henrique998/go-auth/internal/app/request"
	sessionusecases "github.com/henrique998/go-auth/internal/app/usecases/session-usecases"
	"github.com/henrique998/go-auth/internal/infra/database"
	"github.com/henrique998/go-auth/internal/infra/database/repositories"
	"github.com/henrique998/go-auth/internal/infra/providers"
)

func LoginWithCredentialsController(c fiber.Ctx) error {
	ip := c.IP()

	db := database.ConnectToDb()
	defer db.Close()

	repo := repositories.PGAccountsRepository{
		Db: db,
	}
	emailProvider := providers.ResendEmailProvider{ApiKey: os.Getenv("RESEND_API_KEY")}

	usecase := sessionusecases.LoginWithCredentialsUseCase{
		Repo:          &repo,
		EmailProvider: &emailProvider,
	}

	body := c.Body()

	var req request.LoginWithCredentialsRequest
	req.IP = ip

	jsonErr := json.Unmarshal(body, &req)
	if jsonErr != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Internal Server Error.")
	}

	token, err := usecase.Execute(req)
	if err != nil {
		return c.Status(err.GetStatus()).SendString(err.GetMessage())
	}

	return c.SendString(token)
}
