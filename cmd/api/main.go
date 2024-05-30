package main

import (
	"log"

	"github.com/gofiber/fiber/v3"
	"github.com/henrique998/go-auth/internal/infra/endpoints"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}

	app := fiber.New()
	endpoints.SetupEndpoints(app)
	app.Listen(":3333")
}
