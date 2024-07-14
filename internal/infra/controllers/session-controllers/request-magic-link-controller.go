package sessioncontrollers

import (
	"encoding/json"
	"os"

	"github.com/gofiber/fiber/v3"
	sessionusecases "github.com/henrique998/go-auth/internal/app/usecases/session-usecases"
	"github.com/henrique998/go-auth/internal/infra/database"
	"github.com/henrique998/go-auth/internal/infra/database/repositories"
	"github.com/henrique998/go-auth/internal/infra/providers"
)

type Request struct {
	Email string `json:"email"`
}

func RequestMagicLinkController(c fiber.Ctx) error {
	db := database.ConnectToDb()
	defer db.Close()

	repo := repositories.PGAccountsRepository{
		Db: db,
	}
	mlRepo := repositories.PGMagicLinksRepository{
		Db: db,
	}
	emailProvider := providers.ResendEmailProvider{
		ApiKey: os.Getenv("RESEND_API_KEY"),
	}

	usecase := sessionusecases.RequestMagicLinkUseCase{
		Repo:          &repo,
		MLRepo:        &mlRepo,
		EmailProvider: &emailProvider,
	}

	body := c.Body()

	var req Request

	jsonErr := json.Unmarshal(body, &req)
	if jsonErr != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("internal Server Error.")
	}

	err := usecase.Execute(req.Email)
	if err != nil {
		return c.Status(err.GetStatus()).SendString(err.GetMessage())
	}

	return c.SendString("magic link sent successfuly!")
}
