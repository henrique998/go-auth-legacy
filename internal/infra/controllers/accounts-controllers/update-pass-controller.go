package accountscontrollers

import (
	"encoding/json"

	"github.com/gofiber/fiber/v3"
	"github.com/henrique998/go-auth/internal/app/request"
	accountsusecases "github.com/henrique998/go-auth/internal/app/usecases/accounts-usecases"
	"github.com/henrique998/go-auth/internal/infra/database"
	"github.com/henrique998/go-auth/internal/infra/database/repositories"
)

func UpdatePassController(c fiber.Ctx) error {
	db := database.ConnectToDb()
	defer db.Close()

	repo := repositories.PGAccountsRepository{Db: db}
	vtRepo := repositories.PGVerificationCodesRepository{Db: db}

	usecase := accountsusecases.UpdatePassUsecase{
		Repo:   &repo,
		VTRepo: &vtRepo,
	}

	body := c.Body()

	var req request.NewPassRequest

	jsonErr := json.Unmarshal(body, &req)
	if jsonErr != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("internal server error.")
	}

	err := usecase.Execute(req)
	if err != nil {
		return c.Status(err.GetStatus()).SendString(err.GetMessage())
	}

	return c.SendString("password updated successfuly!")
}
