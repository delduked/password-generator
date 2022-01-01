package models

import (
	"github.com/gofiber/fiber/v2"
	"gitlab.com/alienate/password-generator/controllers"
	"gitlab.com/alienate/password-generator/handlers"
	"gitlab.com/alienate/password-generator/schema"
)

func YouPassed(c *fiber.Ctx) error {
	res := schema.Response{
		Status: fiber.StatusOK,
		Error:  nil,
	}

	return handlers.Response(res, c)
}

func Auth(c *fiber.Ctx) error {

	bearer := c.Cookies("authToken")

	_, err := controllers.Verify(bearer)
	if err != nil {
		res := schema.JWT{
			Status: fiber.StatusForbidden,
			Error:  err,
			Valid:  false,
		}
		return handlers.JWTResponse(res, c)
	}

	return c.Next()
}
