package endpoints

import (
	"github.com/gofiber/fiber/v3"
	accountscontrollers "github.com/henrique998/go-auth/internal/infra/controllers/accounts-controllers"
	"github.com/henrique998/go-auth/internal/infra/endpoints/middlewares"
)

func accountsEndpoints(app *fiber.App) {
	app.Post("/accounts", accountscontrollers.CreateAccountController)
	app.Get("/accounts/verify-email", accountscontrollers.VerifyEmailAccountController)
	app.Post("/accounts/send-2fa-code", accountscontrollers.Send2FACodeController)
	app.Post("/accounts/verify-2fa-code", accountscontrollers.Verify2faCodeController)
	app.Post("/accounts/send-new-pass-request", accountscontrollers.SendNewPassRequestController)
	app.Put("/accounts/update-pass", accountscontrollers.UpdatePassController)

	auth := app.Group("/accounts", middlewares.AuthMiddleware())
	auth.Get("/devices", accountscontrollers.GetAccountDevicesController)
}
