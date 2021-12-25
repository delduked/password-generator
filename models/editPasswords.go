package models

import (
	"github.com/gofiber/fiber/v2"
	"gitlab.com/alienate/password-generator/controllers"
	"gitlab.com/alienate/password-generator/handlers"
	"gitlab.com/alienate/password-generator/types"
)

// @Summary Save password
// @Description Save a password to the redis database
// @Accept json
// @Success 200 {object} fiber.StatusOK
// @Failure 400
// @Router /db [post]
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

// @Summary Update password
// @Description Update an existing password in the redis database
// @Accept json
// @Success 200
// @Failure 400
// @Failure 500
// @Router /db [patch]
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

// @Summary Get password
// @Description Get All password in the redis database
// @Accept json
// @Success 200
// @Failure 400
// @Router /db/ [get]
func GetPasswords(c *fiber.Ctx) error {

	allPassword, err := controllers.GetAll()
	if err != nil {
		res := types.Response{
			Status: fiber.StatusInternalServerError,
			Error:  err,
		}
		return handlers.Response(res, c)
	}

	res := types.AllPasswordResponse{
		Status:    fiber.StatusOK,
		Error:     err,
		Passwords: allPassword,
	}
	return handlers.AllPasswordResponse(res, c)
}

// @Summary Get specfic password field
// @Description Get a specfic password field in the redis database
// @Accept json
// @Success 200
// @Failure 400
// @Router /db [get]
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

// @Summary Get specfic password field
// @Description Get a specfic password field in the redis database
// @Accept json
// @Success 200
// @Failure 400
// @Router /db [delete]
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
