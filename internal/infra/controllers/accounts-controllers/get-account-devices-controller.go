package accountscontrollers

import (
	"os"

	"github.com/gofiber/fiber/v3"
	accountsusecases "github.com/henrique998/go-auth/internal/app/usecases/accounts-usecases"
	"github.com/henrique998/go-auth/internal/infra/database"
	"github.com/henrique998/go-auth/internal/infra/database/repositories"
	"github.com/henrique998/go-auth/internal/infra/utils"
)

func GetAccountDevicesController(c fiber.Ctx) error {
	cookie := c.Cookies("goauth:access_token")

	if cookie == "" {
		return c.Status(fiber.StatusBadRequest).SendString("Cookie not found!")
	}

	accountId, err := utils.ParseJWTToken(cookie, os.Getenv("JWT_SECRET"))
	if err != nil {
		return c.Status(err.GetStatus()).SendString(err.GetMessage())
	}

	db := database.ConnectToDb()
	defer db.Close()

	repo := repositories.PGAccountsRepository{
		Db: db,
	}
	devicesRepo := repositories.PGDevicesRepository{
		Db: db,
	}

	usecase := accountsusecases.GetAccountDevicesUseCase{
		Repo:        &repo,
		DevicesRepo: &devicesRepo,
	}

	devices, err := usecase.Execute(accountId)
	if err != nil {
		return c.Status(err.GetStatus()).JSON(err.GetMessage())
	}

	return c.JSON(fiber.Map{
		"devices": devices,
	})
}
