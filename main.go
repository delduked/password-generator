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

	app.Get("/", models.Health)
	app.Get("/healthcheck", models.Health)
	app.Get("/dashboard", monitor.New())

	// generate new password with either JSON or Parameters
	app.Post("/pw", models.GenerateBody)
	app.Get("/pw", models.GenerateParams)

	// get acceess token in order to make requests to the redis database
	app.Post("/signup", models.SignUp)
	app.Post("/signin", models.SignIn)

	// authentication middleware behind redis access
	db := app.Group("/db", models.Auth)
	db.Get("/", models.GetPasswords)
	db.Get("/:key", models.GetKeyedField)
	db.Post("/", models.SavePassword)
	db.Put("/", models.SaveMany)
	db.Patch("/", models.UpdatePassword)
	db.Delete("/:key", models.DeleteKeyedField)

	app.Listen(":8080")
}
