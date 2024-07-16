package middlewares

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber/v3"
	"github.com/henrique998/go-auth/internal/infra/database"
	"github.com/henrique998/go-auth/internal/infra/database/repositories"
	"github.com/henrique998/go-auth/internal/infra/utils"
)

func AuthMiddleware() fiber.Handler {
	return func(c fiber.Ctx) error {
		db := database.ConnectToDb()
		defer db.Close()
		accessTokenStr := c.Cookies("goauth:access_token")

		if accessTokenStr == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Unauthorized - No token provided",
			})
		}

		accountId, err := utils.ParseJWTToken(accessTokenStr, os.Getenv("JWT_SECRET"))
		if err != nil {
			return c.JSON(fiber.Map{
				"error": err.GetMessage(),
			})
		}

		repo := repositories.PGAccountsRepository{
			Db: db,
		}

		account := repo.FindById(accountId)
		fmt.Println(account)

		if account == nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Unauthorized - Account not found!",
			})
		}

		return c.Next()
	}
}
