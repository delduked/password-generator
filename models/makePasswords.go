package models

import (
	"github.com/gofiber/fiber/v2"
	"gitlab.com/alienate/password-generator/controllers"
	"gitlab.com/alienate/password-generator/handlers"
	"gitlab.com/alienate/password-generator/types"
)

// ShowAccount godoc
// @Summary get health check
// @Description get health status of api
// @Produce json
// @Router /healthcheck [get]
func Health(c *fiber.Ctx) error {

	res := types.Response{
		Status: fiber.StatusOK,
		Error:  nil,
	}
	return handlers.Response(res, c)

}

// @Summary Generate new random password
// @Description Generate a new random password with various constraints given by the user.
// @Accept json
// @Produce json
// @Success 200
// @Failure 400
// @Router /pw [get]
func GenerateBody(c *fiber.Ctx) error {
	c.Accepts("application/json")

	body := new(types.NewPasswordRequest)
	if err := c.BodyParser(body); err != nil {
		res := types.Response{
			Status: fiber.StatusBadRequest,
			Error:  err,
		}
		return handlers.Response(res, c)
	}

	password := controllers.GenerateResponse(body)
	res := types.NewPasswordResponse{
		Status:   fiber.StatusOK,
		Error:    nil,
		Password: password,
	}

	return handlers.NewPasswordResponse(res, c)

}

// @Summary Generate new random password
// @Description Generate a new random password with various constraints given by the user.
// @Accept json
// @Produce json
// @Success 200
// @Failure 400
// @Router /pw [get]
func GenerateParams(c *fiber.Ctx) error {
	body := new(types.NewPasswordRequest)
	if err := c.QueryParser(body); err != nil {
		res := types.Response{
			Status: fiber.StatusBadRequest,
			Error:  err,
		}
		return handlers.Response(res, c)
	}

	password := controllers.GenerateResponse(body)
	res := types.NewPasswordResponse{
		Status:   fiber.StatusOK,
		Error:    nil,
		Password: password,
	}

	return handlers.NewPasswordResponse(res, c)

}
