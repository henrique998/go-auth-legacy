package accountscontrollers

import (
	"encoding/json"
	"os"

	"github.com/gofiber/fiber/v3"
	"github.com/henrique998/go-auth/internal/app/request"
	accountsusecases "github.com/henrique998/go-auth/internal/app/usecases/accounts-usecases"
	"github.com/henrique998/go-auth/internal/infra/database"
	"github.com/henrique998/go-auth/internal/infra/database/repositories"
	"github.com/henrique998/go-auth/internal/infra/providers"
)

func CreateAccountController(c fiber.Ctx) error {
	db := database.ConnectToDb()
	defer db.Close()

	repo := repositories.PGAccountsRepository{Db: db}
	vtRepo := repositories.PGVerificationTokensRepository{Db: db}
	emailProvider := providers.ResendEmailProvider{ApiKey: os.Getenv("RESEND_API_KEY")}

	usecase := accountsusecases.CreateAccountUseCase{
		Repo:          &repo,
		VTRepo:        &vtRepo,
		EmailProvider: &emailProvider,
	}

	body := c.Body()

	var req request.CreateAccountRequest

	jsonErr := json.Unmarshal(body, &req)

	if jsonErr != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Internal Server Error.")
	}

	err := usecase.Execute(req)
	if err != nil {
		return c.Status(err.GetStatus()).SendString(err.GetMessage())
	}

	return nil
}
