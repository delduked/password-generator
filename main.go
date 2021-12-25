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
	app.Get("/healthcheck", models.Health)

	pw := app.Group("/pw")

	pw.Post("/", models.GenerateBody)
	pw.Get("/", models.GenerateParams)

	db := app.Group("/db")

	db.Get("/", models.GetPasswords)
	db.Get("/:key", models.GetKeyedField)
	db.Post("/", models.SavePassword)
	db.Patch("/", models.UpdatePassword)
	db.Delete("/:key", models.DeleteKeyedField)

	app.Listen(":8080")
}
