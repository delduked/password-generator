package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"gitlab.com/alienate/password-generator/models"
)

func main() {

	app := fiber.New()

	app.Use(cors.New())

	app.Get("/dashboard", monitor.New())

	// health check
	app.Get("/healthcheck", models.Health)

	// Generate password
	app.Post("/generateBody", models.GenerateBody)
	app.Get("/generateParams", models.GenerateParams)

	// Password endpoints
	app.Get("/passwords", models.GetPasswords)
	app.Get("/password/:key", models.GetKeyedField)
	app.Post("/password", models.SavePassword)
	app.Patch("/password", models.UpdatePassword)
	// app.Delete("/password", routes.DeletePassword)

	app.Listen(":8080")
}
