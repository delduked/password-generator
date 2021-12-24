package models

import (
	"github.com/gofiber/fiber/v2"
	"gitlab.com/alienate/password-generator/controllers"
	"gitlab.com/alienate/password-generator/handlers"
	"gitlab.com/alienate/password-generator/types"
)

func Health(c *fiber.Ctx) error {

	res := types.Response{
		Status: fiber.StatusOK,
		Error:  nil,
	}
	return handlers.Response(res, c)

}

func GenerateBody(c *fiber.Ctx) error {
	c.Accepts("application/json")

	p := new(types.NewPasswordRequest)
	if err := c.BodyParser(p); err != nil {
		res := types.Response{
			Status: fiber.StatusBadRequest,
			Error:  err,
		}
		return handlers.Response(res, c)
	}

	password := controllers.GenerateResponse(p)
	res := types.NewPasswordResponse{
		Status:   fiber.StatusOK,
		Error:    nil,
		Password: password,
	}

	return handlers.NewPasswordResponse(res, c)

}
func GenerateParams(c *fiber.Ctx) error {
	p := new(types.NewPasswordRequest)
	if err := c.QueryParser(p); err != nil {
		res := types.Response{
			Status: fiber.StatusBadRequest,
			Error:  err,
		}
		return handlers.Response(res, c)
	}

	password := controllers.GenerateResponse(p)
	res := types.NewPasswordResponse{
		Status:   fiber.StatusOK,
		Error:    nil,
		Password: password,
	}

	return handlers.NewPasswordResponse(res, c)

}
