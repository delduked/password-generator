package routes

import (
	"github.com/gofiber/fiber/v2"
	"gitlab.com/alienate/password-generator/controllers"
	"gitlab.com/alienate/password-generator/handlers"
	"gitlab.com/alienate/password-generator/inter"
)

func Health(c *fiber.Ctx) error {

	status := inter.Response{
		Status:   fiber.StatusOK,
		Error:    nil,
		Password: "nil",
	}
	return handlers.Response(status, c)

}

func GenerateBody(c *fiber.Ctx) error {
	c.Accepts("application/json")

	var res inter.Response

	p := new(inter.Request)
	if err := c.BodyParser(p); err != nil {
		res = inter.Response{
			Status:   fiber.StatusBadRequest,
			Error:    err,
			Password: "nil",
		}
		return handlers.Response(res, c)
	}

	password := controllers.GenerateResponse(p)
	res = inter.Response{
		Status:   fiber.StatusOK,
		Error:    nil,
		Password: password,
	}

	return handlers.Response(res, c)

}
func GenerateParams(c *fiber.Ctx) error {
	var res inter.Response
	p := new(inter.Request)
	if err := c.QueryParser(p); err != nil {
		res = inter.Response{
			Status:   fiber.StatusBadRequest,
			Error:    err,
			Password: "nil",
		}
		return handlers.Response(res, c)
	}

	password := controllers.GenerateResponse(p)

	res = inter.Response{
		Status:   fiber.StatusOK,
		Error:    nil,
		Password: password,
	}

	return handlers.Response(res, c)

}
