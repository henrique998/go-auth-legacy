package sessioncontrollers

import (
	"encoding/json"
	"time"

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

	accessToken, refreshToken, err := usecase.Execute(req.RefreshToken)
	if err != nil {
		return c.Status(err.GetStatus()).SendString(err.GetMessage())
	}

	accessTokenCookie := fiber.Cookie{
		Name:     "goauth:access_token",
		Value:    accessToken,
		Expires:  time.Now().Add(15 * time.Second),
		HTTPOnly: true,
		Path:     "/",
	}

	refreshTokenCookie := fiber.Cookie{
		Name:     "goauth:refresh_token",
		Value:    refreshToken,
		Expires:  time.Now().Add(time.Hour * 24 * 30),
		HTTPOnly: true,
		Path:     "/",
	}

	c.Cookie(&accessTokenCookie)
	c.Cookie(&refreshTokenCookie)

	return c.JSON(fiber.Map{
		"message": "success",
	})
}
