package models

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"gitlab.com/alienate/password-generator/controllers"
	"gitlab.com/alienate/password-generator/handlers"
	"gitlab.com/alienate/password-generator/types"
)

func SavePassword(c *fiber.Ctx) error {
	c.Accepts("application/json")

	p := new(types.NewPasswordReqSave)
	if err := c.BodyParser(p); err != nil {
		res := types.Response{
			Status: fiber.StatusBadRequest,
			Error:  err,
		}
		fmt.Println(res)
		return handlers.Response(res, c)
	}

	_, err := controllers.Save(p)
	if err != nil {
		res := types.Response{
			Status: fiber.StatusBadRequest,
			Error:  err,
		}
		fmt.Println(res)
		return handlers.Response(res, c)
	}
	res := types.Response{
		Status: fiber.StatusOK,
		Error:  nil,
	}
	fmt.Println(res)
	return handlers.Response(res, c)

}

func UpdatePassword(c *fiber.Ctx) error {
	c.Accepts("application/json")

	p := new(types.SavedFields)
	if err := c.BodyParser(p); err != nil {
		res := types.Response{
			Status: fiber.StatusBadRequest,
			Error:  err,
		}
		return handlers.Response(res, c)
	}

	_, err := controllers.Update(p)
	if err != nil {
		res := types.Response{
			Status: fiber.StatusBadRequest,
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
			Status: fiber.StatusBadRequest,
			Error:  err,
		}
		fmt.Println(res)
		return handlers.Response(res, c)
	}

	res := types.Test{
		Status:    fiber.StatusOK,
		Error:     err,
		Passwords: allPassword,
	}
	fmt.Println(res)
	return handlers.Test(res, c)
}
func GetKeyedField(c *fiber.Ctx) error {
	key := c.Params("key")
	fmt.Println(key)

	KeyedField, err := controllers.GetKeyedPassword(key)
	if err != nil {
		res := types.Response{
			Status: fiber.StatusBadRequest,
			Error:  err,
		}
		fmt.Println(res)
		return handlers.Response(res, c)
	}

	res := types.KeyedResponse{
		Status: fiber.StatusOK,
		Error:  err,
		Fields: KeyedField,
	}
	fmt.Println(res)
	return handlers.KeyedResponse(res, c)
}
