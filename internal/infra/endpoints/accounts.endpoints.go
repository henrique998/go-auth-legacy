package endpoints

import (
	"github.com/gofiber/fiber/v3"
	accountscontrollers "github.com/henrique998/go-auth/internal/infra/controllers/accounts-controllers"
)

func accountsEndpoints(app *fiber.App) {
	app.Post("/accounts", accountscontrollers.CreateAccountController)
	app.Get("/accounts/verify-email", accountscontrollers.VerifyEmailAccountController)
	app.Post("/accounts/send-2fa-code", accountscontrollers.Send2FACodeController)
	app.Post("/accounts/verify-2fa-code", accountscontrollers.Verify2faCodeController)
}
