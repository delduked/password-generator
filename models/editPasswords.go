package models

import (
	"github.com/gofiber/fiber/v2"
	"gitlab.com/alienate/password-generator/controllers"
	"gitlab.com/alienate/password-generator/handlers"
	"gitlab.com/alienate/password-generator/schema"
)

func SavePassword(c *fiber.Ctx) error {
	c.Accepts("application/json")

	body := new(schema.KeyedField)
	err := c.BodyParser(body)
	if err != nil {
		res := schema.Response{
			Status: fiber.StatusBadRequest,
			Error:  err,
		}
		return handlers.Response(res, c)
	}

	username := c.Locals("username").(string)
	savedField, err := controllers.Save(username, body)
	if err != nil {
		res := schema.Response{
			Status: fiber.StatusInternalServerError,
			Error:  err,
		}
		return handlers.Response(res, c)
	}
	res := schema.SavedFieldResponse{
		Status: fiber.StatusOK,
		Error:  nil,
		Field:  savedField,
	}
	return handlers.SavedFieldResponse(res, c)
}

func SaveMany(c *fiber.Ctx) error {
	c.Accepts("application/json")

	body := new([]schema.KeyedField)
	if err := c.BodyParser(body); err != nil {
		res := schema.Response{
			Status: fiber.StatusBadRequest,
			Error:  err,
		}
		return handlers.Response(res, c)
	}

	username := c.Locals("username").(string)
	err := controllers.SaveMany(username, *body)
	if err != nil {
		res := schema.Response{
			Status: fiber.StatusInternalServerError,
			Error:  err,
		}
		return handlers.Response(res, c)
	}

	res := schema.Response{
		Status: fiber.StatusOK,
		Error:  nil,
	}
	return handlers.Response(res, c)
}

func UpdatePassword(c *fiber.Ctx) error {
	c.Accepts("application/json")

	body := new(schema.SavedField)
	if err := c.BodyParser(body); err != nil {
		res := schema.Response{
			Status: fiber.StatusBadRequest,
			Error:  err,
		}
		return handlers.Response(res, c)
	}

	_, err := controllers.Update(body)
	if err != nil {
		res := schema.Response{
			Status: fiber.StatusInternalServerError,
			Error:  err,
		}
		return handlers.Response(res, c)
	}
	res := schema.Response{
		Status: fiber.StatusOK,
		Error:  nil,
	}
	return handlers.Response(res, c)
}

func GetPasswords(c *fiber.Ctx) error {

	username := c.Locals("username").(string)
	savedFields, err := controllers.GetAll(username)
	if err != nil {
		res := schema.Response{
			Status: fiber.StatusInternalServerError,
			Error:  err,
		}
		return handlers.Response(res, c)
	}

	res := schema.SavedFieldsResponse{
		Status: fiber.StatusOK,
		Error:  err,
		Fields: savedFields,
	}
	return handlers.SavedFieldsResponse(res, c)
}

func GetKeyedField(c *fiber.Ctx) error {
	key := c.Params("key")

	username := c.Locals("username").(string)
	KeyedField, err, length := controllers.GetKeyedPassword(username + "::" + key)
	if err != nil {
		res := schema.Response{
			Status: fiber.StatusBadRequest,
			Error:  err,
		}
		return handlers.Response(res, c)
	} else if length < 1 {
		res := schema.Response{
			Status: fiber.StatusNotFound,
			Error:  err,
		}
		return handlers.Response(res, c)
	}

	res := schema.KeyedResponse{
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
		res := schema.Response{
			Status: fiber.StatusBadRequest,
			Error:  err,
		}
		return handlers.Response(res, c)
	}

	res := schema.Response{
		Status: fiber.StatusOK,
		Error:  err,
	}
	return handlers.Response(res, c)
}
