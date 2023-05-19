package middleware

import (
	"InternService/internal/utilities"
	"InternService/internal/utilities/constants"
	"github.com/gofiber/fiber/v2"
	"strings"
)

func Authorize(ctx *fiber.Ctx) error {
	// get authorization header
	rawToken := ctx.Get("Authorization")
	if rawToken == "" {
		return utilities.Response(utilities.ResponseParams{
			Ctx:    ctx,
			Info:   constants.ResponseMessages.MissingToken,
			Status: fiber.StatusUnauthorized,
		})
	}
	trimmedToken := strings.TrimSpace(rawToken)
	if trimmedToken == "" {
		return utilities.Response(utilities.ResponseParams{
			Ctx:    ctx,
			Info:   constants.ResponseMessages.MissingToken,
			Status: fiber.StatusUnauthorized,
		})
	}

	// parse JWT
	claims, parsingError := jwtokens.ParseClaims(trimmedToken)
	if parsingError != nil {
		return utilities.Response(utilities.ResponseParams{
			Ctx:    ctx,
			Info:   constants.ResponseMessages.AccessDenied,
			Status: fiber.StatusUnauthorized,
		})
	}

	// store User ID in Locals so that it can be accessed later and proceed
	ctx.Locals("UserId", claims.UserId)
	return ctx.Next()
}
