package accountscontrollers

import (
	"encoding/json"
	"os"

	"github.com/gofiber/fiber/v3"
	accountsusecases "github.com/henrique998/go-auth/internal/app/usecases/accounts-usecases"
	"github.com/henrique998/go-auth/internal/infra/database"
	"github.com/henrique998/go-auth/internal/infra/database/repositories"
	"github.com/henrique998/go-auth/internal/infra/providers"
)

type Body struct {
	Email string `json:"email"`
}

func SendNewPassRequestController(c fiber.Ctx) error {
	db := database.ConnectToDb()
	defer db.Close()

	repo := repositories.PGAccountsRepository{Db: db}
	vtRepo := repositories.PGVerificationTokensRepository{Db: db}
	emailProvider := providers.ResendEmailProvider{ApiKey: os.Getenv("RESEND_API_KEY")}

	usecase := accountsusecases.SendNewPassRequestUseCase{
		Repo:          &repo,
		VTRepo:        &vtRepo,
		EmailProvider: &emailProvider,
	}

	body := c.Body()

	var req Body

	jsonErr := json.Unmarshal(body, &req)
	if jsonErr != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("internal Server Error.")
	}

	err := usecase.Execute(req.Email)
	if err != nil {
		return c.Status(err.GetStatus()).SendString(err.GetMessage())
	}

	return c.SendString("password request sent successfuly!")
}
