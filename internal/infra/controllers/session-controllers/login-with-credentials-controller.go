package sessioncontrollers

import (
	"encoding/json"
	"os"
	"time"

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
	rtRepo := repositories.PGRefreshTokensRepository{
		Db: db,
	}
	emailProvider := providers.ResendEmailProvider{ApiKey: os.Getenv("RESEND_API_KEY")}

	usecase := sessionusecases.LoginWithCredentialsUseCase{
		Repo:          &repo,
		RTRepo:        &rtRepo,
		EmailProvider: &emailProvider,
	}

	body := c.Body()

	var req request.LoginWithCredentialsRequest
	req.IP = ip

	jsonErr := json.Unmarshal(body, &req)
	if jsonErr != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Internal Server Error.")
	}

	accessToken, refreshToken, err := usecase.Execute(req)
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
