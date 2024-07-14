package endpoints

import (
	"github.com/gofiber/fiber/v3"
	sessioncontrollers "github.com/henrique998/go-auth/internal/infra/controllers/session-controllers"
)

func sessionEndpoints(app *fiber.App) {
	app.Post("/session/login", sessioncontrollers.LoginWithCredentialsController)
	app.Post("/session/logout", sessioncontrollers.LogoutController)
	app.Post("/session/refresh-token", sessioncontrollers.RefreshTokenUseController)

	app.Get("/session/google/redirect", sessioncontrollers.RedirectGoogleLogin)
	app.Get("/session/callback/google", func(c fiber.Ctx) error {
		code := c.Query("code")

		codeMap := map[string]string{
			"code": code,
		}

		return c.JSON(codeMap)
	})
	app.Post("/session/login/google", sessioncontrollers.LoginWithGoogleController)

	app.Post("/session/magic-link/request", sessioncontrollers.RequestMagicLinkController)
	app.Get("/session/login/magic-link", sessioncontrollers.LoginWithMagicLinkController)
}
