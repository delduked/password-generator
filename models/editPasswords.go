package models

import (
	"github.com/gofiber/fiber/v2"
	"gitlab.com/alienate/password-generator/controllers"
	"gitlab.com/alienate/password-generator/handlers"
	"gitlab.com/alienate/password-generator/types"
)

func SavePassword(c *fiber.Ctx) error {
	c.Accepts("application/json")

	body := new(types.NewPasswordReqSave)
	if err := c.BodyParser(body); err != nil {
		res := types.Response{
			Status: fiber.StatusBadRequest,
			Error:  err,
		}
		return handlers.Response(res, c)
	}

	_, err := controllers.Save(body)
	if err != nil {
		res := types.Response{
			Status: fiber.StatusInternalServerError,
			Error:  err,
		}
		return handlers.Response(res, c)
	}
	res := types.Response{
		Status: fiber.StatusOK,
		Error:  nil,
	}
	return handlers.Response(res, c)

}

func UpdatePassword(c *fiber.Ctx) error {
	c.Accepts("application/json")

	body := new(types.SavedFields)
	if err := c.BodyParser(body); err != nil {
		res := types.Response{
			Status: fiber.StatusBadRequest,
			Error:  err,
		}
		return handlers.Response(res, c)
	}

	_, err := controllers.Update(body)
	if err != nil {
		res := types.Response{
			Status: fiber.StatusInternalServerError,
			Error:  err,
		}
		return handlers.Response(res, c)
	}
	res := types.Response{
		Status: fiber.StatusOK,
		Error:  nil,
	}
	return handlers.Response(res, c)
}

func GetPasswords(c *fiber.Ctx) error {

	allPassword, err := controllers.GetAll()
	if err != nil {
		res := types.Response{
			Status: fiber.StatusInternalServerError,
			Error:  err,
		}
		return handlers.Response(res, c)
	}

	res := types.Test{
		Status:    fiber.StatusOK,
		Error:     err,
		Passwords: allPassword,
	}
	return handlers.Test(res, c)
}
func GetKeyedField(c *fiber.Ctx) error {
	key := c.Params("key")

	KeyedField, err := controllers.GetKeyedPassword(key)
	if err != nil {
		res := types.Response{
			Status: fiber.StatusBadRequest,
			Error:  err,
		}
		return handlers.Response(res, c)
	}

	res := types.KeyedResponse{
		Status: fiber.StatusOK,
		Error:  err,
		Fields: KeyedField,
	}
	return handlers.KeyedResponse(res, c)
}
func DeleteKeyedField(c *fiber.Ctx) error {
	key := c.Params("key")

	_, err := controllers.Delete(key)
	if err != nil {
		res := types.Response{
			Status: fiber.StatusBadRequest,
			Error:  err,
		}
		return handlers.Response(res, c)
	}

	res := types.Response{
		Status: fiber.StatusOK,
		Error:  err,
	}
	return handlers.Response(res, c)
}
