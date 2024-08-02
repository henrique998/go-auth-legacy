package sessioncontrollers

import (
	"os"
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/henrique998/go-auth/internal/app/request"
	sessionusecases "github.com/henrique998/go-auth/internal/app/usecases/session-usecases"
	"github.com/henrique998/go-auth/internal/infra/database"
	"github.com/henrique998/go-auth/internal/infra/database/repositories"
	"github.com/henrique998/go-auth/internal/infra/providers"
)

func LoginWithMagicLinkController(c fiber.Ctx) error {
	code := c.Query("code")

	if code == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Code is required"})
	}

	db := database.ConnectToDb()
	defer db.Close()

	ip := c.IP()
	userAgent := c.Get("User-Agent")

	repo := repositories.PGAccountsRepository{
		Db: db,
	}
	mlRepo := repositories.PGMagicLinksRepository{
		Db: db,
	}
	devicesRepo := repositories.PGDevicesRepository{
		Db: db,
	}
	rtRepo := repositories.PGRefreshTokensRepository{
		Db: db,
	}
	atProvider := providers.AuthTokensProvider{
		RTRepo: &rtRepo,
	}
	emailProvider := providers.ResendEmailProvider{
		ApiKey: os.Getenv("RESEND_API_KEY"),
	}
	glProvider := providers.IPStackGeoLocationProvider{
		APiKey: os.Getenv("IPSTACK_API_KEY"),
	}

	usecase := sessionusecases.LoginWithMagicLinkUseCase{
		Repo:          &repo,
		DevicesRepo:   &devicesRepo,
		MLRepo:        &mlRepo,
		ATProvider:    &atProvider,
		EmailProvider: &emailProvider,
		GLProvider:    &glProvider,
	}

	var req request.LoginWithMagicLinkRequest
	req.Code = code
	req.IP = ip
	req.UserAgent = userAgent

	accessToken, refreshToken, err := usecase.Execute(req)
	if err != nil {
		return c.Status(err.GetStatus()).SendString(err.GetMessage())
	}

	accessTokenCookie := fiber.Cookie{
		Name:     "goauth:access_token",
		Value:    accessToken,
		Expires:  time.Now().Add(15 * time.Minute),
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
