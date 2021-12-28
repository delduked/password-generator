package models

import (
	"github.com/gofiber/fiber/v2"
	"gitlab.com/alienate/password-generator/controllers/jwt"
	"gitlab.com/alienate/password-generator/handlers"
	"gitlab.com/alienate/password-generator/types"
)

func AuthMiddleware(c *fiber.Ctx) error {

	jwt, err := jwt.CheckForJwt()
	if err != nil {
		res := types.JWT{
			Status: fiber.StatusForbidden,
			Error:  err,
			Valid:  false,
		}
		return handlers.JWTResponse(res, c)
	}
	res := types.JWT{
		Status: fiber.StatusOK,
		Error:  nil,
		Valid:  true,
		Bearer: jwt,
	}

	return handlers.JWTResponse(res, c)
}
