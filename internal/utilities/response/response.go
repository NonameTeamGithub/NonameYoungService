package response

import (
	"InternService/internal/utilities"
	"InternService/internal/utilities/constants"
	"github.com/gofiber/fiber/v2"
)

func Response(params ResponseParams) error {
	data := params.Data
	info := params.Info
	status := params.Status
	if info == "" {
		info = constants.ResponseMessages.Ok
	}
	if status == 0 {
		status = fiber.StatusOK
	}

	// caclulate request latency
	initial := params.Ctx.Context().Time()
	latency := utilities.MakeTimestamp() - (initial.UnixNano() / 1e6)

	// create a response map
	responseMap := fiber.Map{
		"datetime": utilities.MakeTimestamp(),
		"info":     info,
		"latency":  latency,
		"request":  params.Ctx.OriginalURL() + " [" + params.Ctx.Method() + "]",
		"status":   status,
	}

	if data != nil {
		responseMap["data"] = data
	}

	return params.Ctx.Status(params.Status).JSON(responseMap)
}
