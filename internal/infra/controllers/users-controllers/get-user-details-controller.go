package userscontrollers

import (
	"github.com/gofiber/fiber/v3"
	usersusecases "github.com/henrique998/go-auth/internal/app/usecases/users-usecases"
	"github.com/henrique998/go-auth/internal/infra/database"
	"github.com/henrique998/go-auth/internal/infra/database/repositories"
	viewmodels "github.com/henrique998/go-auth/internal/infra/view-models"
)

func GetUserDetailsController(c fiber.Ctx) error {
	db := database.ConnectToDb()
	defer db.Close()

	repo := repositories.PGUsersRepository{Db: db}
	usecase := usersusecases.GetUserDetailsUseCase{Repo: &repo}

	user, err := usecase.Execute("henrique@gmail.com")
	if err != nil {
		return c.Status(err.GetStatus()).SendString(err.GetMessage())
	}

	uv := viewmodels.ToUserHttp(user)

	return c.JSON(uv)
}
