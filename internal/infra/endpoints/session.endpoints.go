package endpoints

import (
	"github.com/gofiber/fiber/v3"
	sessioncontrollers "github.com/henrique998/go-auth/internal/infra/controllers/session-controllers"
)

func sessionEndpoints(app *fiber.App) {
	app.Post("/login", sessioncontrollers.LoginWithCredentialsController)
	app.Post("/refresh-token", sessioncontrollers.RefreshTokenUseController)
}
