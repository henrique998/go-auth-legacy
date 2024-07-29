package accountscontrollers

import (
	"github.com/gofiber/fiber/v3"
	accountsusecases "github.com/henrique998/go-auth/internal/app/usecases/accounts-usecases"
	"github.com/henrique998/go-auth/internal/infra/database"
	"github.com/henrique998/go-auth/internal/infra/database/repositories"
)

func VerifyEmailAccountController(c fiber.Ctx) error {
	param := c.Query("token")

	if param == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Missing query parameter 'param'",
		})
	}

	db := database.ConnectToDb()
	defer db.Close()

	repo := repositories.PGAccountsRepository{
		Db: db,
	}
	vtRepo := repositories.PGVerificationCodesRepository{
		Db: db,
	}

	usecase := accountsusecases.VerifyEmailUseCase{
		Repo:   &repo,
		VTRepo: &vtRepo,
	}

	err := usecase.Execute(param)
	if err != nil {
		return c.Status(err.GetStatus()).SendString(err.GetMessage())
	}

	return c.SendString("Email verified successfully!")
}
