package models

import (
	"github.com/gofiber/fiber/v2"
	"gitlab.com/alienate/password-generator/controllers"
	"gitlab.com/alienate/password-generator/handlers"
	"gitlab.com/alienate/password-generator/schema"
)

func Health(c *fiber.Ctx) error {

	res := schema.Response{
		Status: fiber.StatusOK,
		Error:  nil,
	}
	return handlers.Response(res, c)

}

func GenerateBody(c *fiber.Ctx) error {
	c.Accepts("application/json")

	body := new(schema.NewPasswordRequest)
	if err := c.BodyParser(body); err != nil {
		res := schema.Response{
			Status: fiber.StatusBadRequest,
			Error:  err,
		}
		return handlers.Response(res, c)
	}

	password := controllers.GenerateResponse(body)
	res := schema.NewPasswordResponse{
		Status:   fiber.StatusOK,
		Error:    nil,
		Password: password,
	}

	return handlers.NewPasswordResponse(res, c)

}

func GenerateParams(c *fiber.Ctx) error {
	body := new(schema.NewPasswordRequest)
	if err := c.QueryParser(body); err != nil {
		res := schema.Response{
			Status: fiber.StatusBadRequest,
			Error:  err,
		}
		return handlers.Response(res, c)
	}

	password := controllers.GenerateResponse(body)
	res := schema.NewPasswordResponse{
		Status:   fiber.StatusOK,
		Error:    nil,
		Password: password,
	}

	return handlers.NewPasswordResponse(res, c)

}
