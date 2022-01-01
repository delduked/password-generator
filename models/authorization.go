package models

import (
	"github.com/gofiber/fiber/v2"
	"gitlab.com/alienate/password-generator/controllers"
	"gitlab.com/alienate/password-generator/handlers"
	"gitlab.com/alienate/password-generator/schema"
)

func GenerateNewToken(c *fiber.Ctx) error {
	c.Accepts("application/json")

	// check if body contains required fields
	body := new(schema.UserAccount)
	if err := c.BodyParser(body); err != nil {
		res := schema.JWT{
			Status: fiber.StatusForbidden,
			Error:  err,
			Valid:  false,
		}
		return handlers.JWTResponse(res, c)
	}

	// check if user credentials passed are accurate
	err := controllers.CheckCredentials(body)
	if err != nil {
		res := schema.JWT{
			Status: fiber.StatusUnauthorized,
			Error:  err,
			Valid:  false,
		}
		return handlers.JWTResponse(res, c)
	}

	// Generate new token for request
	jwt, err := controllers.GenerateNewToken()
	if err != nil {
		res := schema.JWT{
			Status: fiber.StatusForbidden,
			Error:  err,
			Valid:  false,
		}
		return handlers.JWTResponse(res, c)
	}

	// Generate new cookie for the clients request
	cookie := new(fiber.Cookie)
	cookie.Name = "authToken"
	cookie.Value = jwt

	// Set the JWT inside the newly created cookie
	c.Cookie(cookie)
	res := schema.JWT{
		Status: fiber.StatusOK,
		Error:  nil,
		Valid:  true,
		Bearer: jwt,
	}
	return handlers.JWTResponse(res, c)
}
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
