package models

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"gitlab.com/alienate/password-generator/controllers"
	"gitlab.com/alienate/password-generator/handlers"
	"gitlab.com/alienate/password-generator/types"
)

func GenerateNewToken(c *fiber.Ctx) error {

	jwt, err := controllers.NewToken()
	if err != nil {
		res := types.JWT{
			Status: fiber.StatusForbidden,
			Error:  err,
			Valid:  false,
		}
		return handlers.JWTResponse(res, c)
	}
	cookie := new(fiber.Cookie)
	cookie.Name = "authToken"
	cookie.Value = jwt

	c.Cookie(cookie)
	res := types.JWT{
		Status: fiber.StatusOK,
		Error:  nil,
		Valid:  true,
		Bearer: jwt,
	}

	return handlers.JWTResponse(res, c)
}
func YouPassed(c *fiber.Ctx) error {
	res := types.Response{
		Status: fiber.StatusOK,
		Error:  nil,
	}

	return handlers.Response(res, c)
}

func Auth(c *fiber.Ctx) error {

	bearer := c.Cookies("authToken")
	if bearer == "" {
		res := types.JWT{
			Status: fiber.StatusBadRequest,
			Error:  fmt.Errorf("No bearer token present"),
			Valid:  false,
		}
		return handlers.JWTResponse(res, c)
	}

	_, err := controllers.Verify(bearer)
	if err != nil {
		res := types.JWT{
			Status: fiber.StatusForbidden,
			Error:  err,
			Valid:  false,
		}
		return handlers.JWTResponse(res, c)
	}

	return c.Next()
}
