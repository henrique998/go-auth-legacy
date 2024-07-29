package accountscontrollers

import (
	"encoding/json"
	"os"

	"github.com/gofiber/fiber/v3"
	"github.com/henrique998/go-auth/internal/app/request"
	accountsusecases "github.com/henrique998/go-auth/internal/app/usecases/accounts-usecases"
	"github.com/henrique998/go-auth/internal/infra/database"
	"github.com/henrique998/go-auth/internal/infra/database/repositories"
	"github.com/henrique998/go-auth/internal/infra/utils"
)

func Verify2faCodeController(c fiber.Ctx) error {
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

	repo := repositories.PGAccountsRepository{Db: db}
	vtRepo := repositories.PGVerificationCodesRepository{Db: db}

	usecase := accountsusecases.Verify2faCodeUseCase{
		Repo:   &repo,
		VTRepo: &vtRepo,
	}

	body := c.Body()

	var req request.Verify2faRequest

	jsonErr := json.Unmarshal(body, &req)
	if jsonErr != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Internal Server Error.")
	}

	req.AccountId = accountId

	err = usecase.Execute(req)
	if err != nil {
		return c.Status(err.GetStatus()).SendString(err.GetMessage())
	}

	return c.SendString("Two factor authentication done succesfully!")
}
