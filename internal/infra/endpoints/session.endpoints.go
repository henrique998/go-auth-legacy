package endpoints

import (
	"github.com/gofiber/fiber/v3"
	sessioncontrollers "github.com/henrique998/go-auth/internal/infra/controllers/session-controllers"
)

func sessionEndpoints(app *fiber.App) {
	app.Post("/session/login", sessioncontrollers.LoginWithCredentialsController)
	app.Post("/session/logout", sessioncontrollers.LogoutController)
	app.Post("/session/refresh-token", sessioncontrollers.RefreshTokenUseController)
}
