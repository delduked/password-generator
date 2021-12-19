package routes

import (
	"encoding/json"

	"github.com/gofiber/fiber/v2"
	"gitlab.com/alienate/password-generator/controllers"
	"gitlab.com/alienate/password-generator/inter"
)

func Health(c *fiber.Ctx) error {

	status := inter.StatusResponse{fiber.StatusOK, nil}
	return statusJson(status, c)

}

func Generate(c *fiber.Ctx) error {
	c.Accepts("application/json")
	// Post request with params
	// requires length int, specials bool, numbers bool, lowerCases bool, upperCases bool
	p := new(inter.Request)
	if err := c.BodyParser(p); err != nil {
		return err
	}

	password := controllers.GenerateResponse(p)
	test := inter.Response{password}

	return responseJson(test, c)

}

func statusJson(asdf inter.StatusResponse, ctx *fiber.Ctx) error {
	writer := ctx.Type("json", "utf-8").Response().BodyWriter()
	return json.NewEncoder(writer).Encode(asdf)
}
func responseJson(asdf inter.Response, ctx *fiber.Ctx) error {
	writer := ctx.Type("json", "utf-8").Response().BodyWriter()
	return json.NewEncoder(writer).Encode(asdf)
}
