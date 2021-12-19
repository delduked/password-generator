package main

import (
	"github.com/gofiber/fiber/v2"
	"gitlab.com/alienate/password-generator/routes"
)

func main() {

	app := fiber.New()

	app.Get("/healthcheck", routes.Health)
	app.Post("/generateBody", routes.GenerateBody)
	app.Get("/generateParams", routes.GenerateParams)

	app.Listen(":8080")
}
