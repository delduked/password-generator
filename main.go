package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"gitlab.com/alienate/password-generator/models"
)

// @title Store and Generate random passwords
// @version 1.0
// @description GO API to save and generate passwords in a redisd database.
// @contact.name Nate Del Duca
// @contact.email nate@nated.ca
// @host localhost:8080
// @BasePath /
func main() {

	app := fiber.New()
	app.Use(cors.New())

	//app.Get("/swagger/*", swagger.Handler)

	app.Get("/dashboard", monitor.New())
	app.Get("/healthcheck", models.Health)

	//auth := app.Group("/auth")
	//auth.Post("/register", models.Register)
	//auth.Post("/signin", models.Signin)

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
