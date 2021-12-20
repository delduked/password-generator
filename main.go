package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"gitlab.com/alienate/password-generator/routes"
)

func main() {

	app := fiber.New()

	app.Use(cors.New())

	app.Get("/healthcheck", routes.Health)
	app.Post("/generateBody", routes.GenerateBody)
	app.Get("/generateParams", routes.GenerateParams)

	app.Listen(":8080")
}
