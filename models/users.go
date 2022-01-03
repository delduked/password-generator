package models

import (
	"github.com/gofiber/fiber/v2"
	"gitlab.com/alienate/password-generator/controllers"
	"gitlab.com/alienate/password-generator/handlers"
	"gitlab.com/alienate/password-generator/schema"
)

func SignUp(c *fiber.Ctx) error {
	c.Accepts("application/json")

	// check if body contains required fields
	newUser := new(schema.SignUp)
	if err := c.BodyParser(newUser); err != nil {
		res := schema.JWT{
			Status: fiber.StatusForbidden,
			Error:  err,
			Valid:  false,
		}
		return handlers.JWTResponse(res, c)
	}

	// check if user credentials passed are accurate
	err := controllers.CheckSecret(newUser)
	if err != nil {
		res := schema.JWT{
			Status: fiber.StatusUnauthorized,
			Error:  err,
			Valid:  false,
		}
		return handlers.JWTResponse(res, c)
	}

	// Save new User
	err = controllers.SaveUser(newUser)
	if err != nil {
		res := schema.JWT{
			Status: fiber.StatusInternalServerError,
			Error:  err,
			Valid:  false,
		}
		return handlers.JWTResponse(res, c)
	}

	// Generate new token for request
	jwt, err := controllers.NewTokenWithUserName(newUser)
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
func SignIn(c *fiber.Ctx) error {
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
	accurate, err := controllers.CheckIfOldUserExists(body)
	if err != nil || !accurate {
		res := schema.JWT{
			Status: fiber.StatusUnauthorized,
			Error:  err,
			Valid:  false,
		}
		return handlers.JWTResponse(res, c)
	}

	user := new(schema.SignUp)
	user.Username = body.Username
	user.Password = body.Password

	// Generate new token for requestu
	jwt, err := controllers.NewTokenWithUserName(user)
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
