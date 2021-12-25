package handlers

import (
	"encoding/json"

	"github.com/gofiber/fiber/v2"
	"gitlab.com/alienate/password-generator/types"
)

func Response(res types.Response, ctx *fiber.Ctx) error {
	writer := ctx.Type("json", "utf-8").Response().BodyWriter()
	return json.NewEncoder(writer).Encode(res)
}
func NewPasswordResponse(res types.NewPasswordResponse, ctx *fiber.Ctx) error {
	writer := ctx.Type("json", "utf-8").Response().BodyWriter()
	return json.NewEncoder(writer).Encode(res)
}

func AllPasswordResponse(res types.AllPasswordResponse, ctx *fiber.Ctx) error {
	writer := ctx.Type("json", "utf-8").Response().BodyWriter()
	return json.NewEncoder(writer).Encode(res)
}
func Test(res types.AllPasswordResponse, ctx *fiber.Ctx) error {
	writer := ctx.Type("json", "utf-8").Response().BodyWriter()
	return json.NewEncoder(writer).Encode(res)
}
func KeyedResponse(res types.KeyedResponse, ctx *fiber.Ctx) error {
	writer := ctx.Type("json", "utf-8").Response().BodyWriter()
	return json.NewEncoder(writer).Encode(res)
}
