package sessioncontrollers

import (
	"encoding/json"

	"github.com/gofiber/fiber/v3"
	"github.com/henrique998/go-auth/internal/app/request"
	sessionusecases "github.com/henrique998/go-auth/internal/app/usecases/session-usecases"
	"github.com/henrique998/go-auth/internal/infra/database"
	"github.com/henrique998/go-auth/internal/infra/database/repositories"
)

func RefreshTokenUseController(c fiber.Ctx) error {
	db := database.ConnectToDb()
	defer db.Close()

	repo := repositories.PGRefreshTokensRepository{
		Db: db,
	}

	usecase := sessionusecases.RefreshTokenUseCase{
		Repo: &repo,
	}

	body := c.Body()

	var req request.RefreshTokenRequest

	jsonErr := json.Unmarshal(body, &req)
	if jsonErr != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Internal Server Error.")
	}

	res, err := usecase.Execute(req.RefreshToken)
	if err != nil {
		return c.Status(err.GetStatus()).SendString(err.GetMessage())
	}

	return c.JSON(res)
}
