package models

import (
	"github.com/gofiber/fiber/v2"
	"gitlab.com/alienate/password-generator/controllers"
	"gitlab.com/alienate/password-generator/handlers"
	"gitlab.com/alienate/password-generator/inter"
)

func Health(c *fiber.Ctx) error {

	status := inter.NewPasswordResponse{
		Status:   fiber.StatusOK,
		Error:    nil,
		Password: "nil",
	}
	return handlers.Response(status, c)

}

func GenerateBody(c *fiber.Ctx) error {
	c.Accepts("application/json")

	var res inter.NewPasswordResponse
	p := new(inter.NewPasswordRequest)
	if err := c.BodyParser(p); err != nil {
		res = inter.NewPasswordResponse{
			Status:   fiber.StatusBadRequest,
			Error:    err,
			Password: "nil",
		}
		return handlers.Response(res, c)
	}

	// asdf := string(c.Body())
	// fmt.Println(asdf)
	password := controllers.GenerateResponse(p)
	res = inter.NewPasswordResponse{
		Status:   fiber.StatusOK,
		Error:    nil,
		Password: password,
	}

	return handlers.Response(res, c)

}
func GenerateParams(c *fiber.Ctx) error {
	var res inter.NewPasswordResponse
	p := new(inter.NewPasswordRequest)
	if err := c.QueryParser(p); err != nil {
		res = inter.NewPasswordResponse{
			Status:   fiber.StatusBadRequest,
			Error:    err,
			Password: "nil",
		}
		return handlers.Response(res, c)
	}

	password := controllers.GenerateResponse(p)

	res = inter.NewPasswordResponse{
		Status:   fiber.StatusOK,
		Error:    nil,
		Password: password,
	}

	return handlers.Response(res, c)

}
