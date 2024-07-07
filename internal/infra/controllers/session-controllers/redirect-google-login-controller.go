package sessioncontrollers

import (
	"github.com/gofiber/fiber/v3"
	googleoauth "github.com/henrique998/go-auth/internal/configs/google-oauth"
)

func RedirectGoogleLogin(c fiber.Ctx) error {
	googleOauthConfig := googleoauth.GetGoogleOauthConfig()

	url := googleOauthConfig.AuthCodeURL("random")

	return c.SendString(url)
}
