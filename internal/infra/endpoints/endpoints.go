package endpoints

import (
	"github.com/gofiber/fiber/v3"
)

func SetupEndpoints(app *fiber.App) {
	accountsEndpoints(app)
}
