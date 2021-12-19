package main

import (
	"github.com/gofiber/fiber/v2"
	"gitlab.com/alienate/password-generator/routes"
)

func main() {

	app := fiber.New()

	app.Get("/healthcheck", routes.Health)
	app.Post("/generate", routes.Generate)

	app.Listen(":8080")
}
