package endpoints

import (
	"github.com/gofiber/fiber/v3"
	userscontrollers "github.com/henrique998/go-auth/internal/infra/controllers/users-controllers"
)

func usersEndpoints(app *fiber.App) {
	app.Post("/users", userscontrollers.AddUserController)
	app.Get("/me", userscontrollers.GetUserDetailsController)
}
