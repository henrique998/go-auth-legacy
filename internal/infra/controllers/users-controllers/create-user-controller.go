package userscontrollers

import (
	"encoding/json"

	"github.com/gofiber/fiber/v3"
	"github.com/henrique998/go-auth/internal/app/request"
	usersusecases "github.com/henrique998/go-auth/internal/app/usecases/users-usecases"
	"github.com/henrique998/go-auth/internal/infra/database"
	"github.com/henrique998/go-auth/internal/infra/database/repositories"
)

func AddUserController(c fiber.Ctx) error {
	db := database.ConnectToDb()
	defer db.Close()

	repo := repositories.PGUsersRepository{Db: db}
	usecase := usersusecases.CreateUserUseCase{Repo: &repo}

	body := c.Body()

	var req request.CreateUserRequest

	jsonErr := json.Unmarshal(body, &req)

	if jsonErr != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Internal Server Error.")
	}

	err := usecase.Execute(req)
	if err != nil {
		return c.Status(err.GetStatus()).SendString(err.GetMessage())
	}

	return nil
}
