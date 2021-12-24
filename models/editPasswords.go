package models

import (
	"github.com/gofiber/fiber/v2"
	"gitlab.com/alienate/password-generator/controllers"
	"gitlab.com/alienate/password-generator/handlers"
	"gitlab.com/alienate/password-generator/types"
)

func SavePassword(c *fiber.Ctx) error {
	c.Accepts("application/json")

	var res types.NewPasswordResponse
	p := new(types.NewPasswordReqSave)

	if err := c.BodyParser(p); err != nil {
		res = types.NewPasswordResponse{
			Status:   fiber.StatusBadRequest,
			Error:    err,
			Password: "nil",
		}
		return handlers.Response(res, c)
	}

	_, err := controllers.Save(p)
	if err != nil {
		status := types.SaveResponse{
			Status: fiber.StatusBadRequest,
			Error:  err,
		}
		return handlers.SaveResponse(status, c)
	}
	status := types.SaveResponse{
		Status: fiber.StatusOK,
		Error:  nil,
	}
	return handlers.SaveResponse(status, c)

}

func UpdatePassword(c *fiber.Ctx) error {
	c.Accepts("application/json")

	var res types.NewPasswordResponse
	p := new(types.SavedFields)

	if err := c.BodyParser(p); err != nil {
		res = types.NewPasswordResponse{
			Status:   fiber.StatusBadRequest,
			Error:    err,
			Password: "nil",
		}
		return handlers.Response(res, c)
	}

	_, err := controllers.Update(p)
	if err != nil {
		status := types.SaveResponse{
			Status: fiber.StatusBadRequest,
			Error:  err,
		}
		return handlers.SaveResponse(status, c)
	}
	status := types.SaveResponse{
		Status: fiber.StatusOK,
		Error:  nil,
	}
	return handlers.SaveResponse(status, c)
}

func GetPasswords(c *fiber.Ctx) error {
	c.Accepts("application/json")

	var res types.NewPasswordResponse
	p := new(types.SavedFields)

	if err := c.BodyParser(p); err != nil {
		res = types.NewPasswordResponse{
			Status:   fiber.StatusBadRequest,
			Error:    err,
			Password: "nil",
		}
		return handlers.Response(res, c)
	}

	allPassword, err := controllers.GetAll()
	if err != nil {
		status := types.SaveResponse{
			Status: fiber.StatusBadRequest,
			Error:  err,
		}
		return handlers.SaveResponse(status, c)
	}

	status := types.AllPasswordResponse{
		Status:    fiber.StatusOK,
		Passwords: allPassword,
	}
	return handlers.AllPasswordResponse(status, c)
}
