package sessioncontrollers

import (
	"time"

	"github.com/gofiber/fiber/v3"
)

func LogoutController(c fiber.Ctx) error {
	accessTokenCookie := fiber.Cookie{
		Name:     "goauth:access_token",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HTTPOnly: true,
		Path:     "/",
	}

	refreshTokenCookie := fiber.Cookie{
		Name:     "goauth:refresh_token",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HTTPOnly: true,
		Path:     "/",
	}

	c.Cookie(&accessTokenCookie)
	c.Cookie(&refreshTokenCookie)

	return c.JSON(fiber.Map{
		"message": "success",
	})
}
