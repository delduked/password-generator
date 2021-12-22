package models

import (
	"github.com/gofiber/fiber/v2"
	"gitlab.com/alienate/password-generator/controllers"
	"gitlab.com/alienate/password-generator/handlers"
	"gitlab.com/alienate/password-generator/inter"
)

func SavePassword(c *fiber.Ctx) error {
	c.Accepts("application/json")

	var res inter.NewPasswordResponse
	p := new(inter.NewPasswordReqSave)

	if err := c.BodyParser(p); err != nil {
		res = inter.NewPasswordResponse{
			Status:   fiber.StatusBadRequest,
			Error:    err,
			Password: "nil",
		}
		return handlers.Response(res, c)
	}

	_, err := controllers.Save(p)
	if err != nil {
		status := inter.SaveResponse{
			Status: fiber.StatusBadRequest,
			Error:  err,
		}
		return handlers.SaveResponse(status, c)
	}
	status := inter.SaveResponse{
		Status: fiber.StatusOK,
		Error:  nil,
	}
	return handlers.SaveResponse(status, c)

}

func UpdatePassword(c *fiber.Ctx) error {
	c.Accepts("application/json")

	var res inter.NewPasswordResponse
	p := new(inter.SavedFields)

	if err := c.BodyParser(p); err != nil {
		res = inter.NewPasswordResponse{
			Status:   fiber.StatusBadRequest,
			Error:    err,
			Password: "nil",
		}
		return handlers.Response(res, c)
	}

	_, err := controllers.Update(p)
	if err != nil {
		status := inter.SaveResponse{
			Status: fiber.StatusBadRequest,
			Error:  err,
		}
		return handlers.SaveResponse(status, c)
	}
	status := inter.SaveResponse{
		Status: fiber.StatusOK,
		Error:  nil,
	}
	return handlers.SaveResponse(status, c)
}

// func DeletePassword(ctx *fiber.Ctx) error {

// }
