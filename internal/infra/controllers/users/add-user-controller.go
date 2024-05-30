package users

import (
	"encoding/json"

	"github.com/gofiber/fiber/v3"
	"github.com/henrique998/go-setup/internal/app/usecases"
	"github.com/henrique998/go-setup/internal/infra/database"
	"github.com/henrique998/go-setup/internal/infra/database/repositories"
)

type RequestData struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func AddUserController(c fiber.Ctx) error {
	db := database.ConnectToDb()
	defer db.Close()

	repo := repositories.PGUsersRepository{Db: db}
	usecase := usecases.CreateUserUseCase{Repo: &repo}

	body := c.Body()

	var req RequestData

	jsonErr := json.Unmarshal(body, &req)

	if jsonErr != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Error trying to decode JSON.")
	}

	err := usecase.Execute(req.Name, req.Email, req.Password)
	if err != nil {
		return c.Status(err.GetStatus()).SendString(err.GetMessage())
	}

	return nil
}
