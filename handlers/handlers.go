package handlers

import (
	"encoding/json"

	"github.com/gofiber/fiber/v2"
	"gitlab.com/alienate/password-generator/schema"
)

func Response(res schema.Response, ctx *fiber.Ctx) error {
	writer := ctx.Type("json", "utf-8").Response().BodyWriter()
	return json.NewEncoder(writer).Encode(res)
}
func NewPasswordResponse(res schema.NewPasswordResponse, ctx *fiber.Ctx) error {
	writer := ctx.Type("json", "utf-8").Response().BodyWriter()
	return json.NewEncoder(writer).Encode(res)
}
func SavedFieldsResponse(res schema.SavedFieldsResponse, ctx *fiber.Ctx) error {
	writer := ctx.Type("json", "utf-8").Response().BodyWriter()
	return json.NewEncoder(writer).Encode(res)
}
func KeyedResponse(res schema.KeyedResponse, ctx *fiber.Ctx) error {
	writer := ctx.Type("json", "utf-8").Response().BodyWriter()
	return json.NewEncoder(writer).Encode(res)
}
func SavedFieldResponse(res schema.SavedFieldResponse, ctx *fiber.Ctx) error {
	writer := ctx.Type("json", "utf-8").Response().BodyWriter()
	return json.NewEncoder(writer).Encode(res)
}
func JWTResponse(res schema.JWT, ctx *fiber.Ctx) error {
	writer := ctx.Type("json", "utf-8").Response().BodyWriter()
	return json.NewEncoder(writer).Encode(res)
}
