package endpoints

import (
	"github.com/gofiber/fiber/v3"
	"github.com/henrique998/go-setup/internal/infra/controllers/users"
	"github.com/henrique998/go-setup/internal/infra/database"
)

func usersEndpoints(app *fiber.App) {
	database.ConnectToDb()
	// // usersRepo := repositories.NewUsersRepository(db)
	// // userscontroller := controllers.NewUsersController(usersRepo)

	app.Get("/users", users.GetUsersController)
	app.Post("/users", users.AddUserController)
}
