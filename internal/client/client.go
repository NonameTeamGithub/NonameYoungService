package client

import (
	"InternService/internal/client/handlers"
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"os"
	"time"
)

func NewClient(ctx context.Context) *fiber.App {
	app := fiber.New(fiber.Config{
		Prefork:                      false,
		ServerHeader:                 "",
		StrictRouting:                false,
		CaseSensitive:                false,
		Immutable:                    false,
		UnescapePath:                 false,
		ETag:                         false,
		BodyLimit:                    0,
		Concurrency:                  0,
		Views:                        nil,
		ViewsLayout:                  "",
		PassLocalsToViews:            false,
		ReadTimeout:                  0,
		WriteTimeout:                 0,
		IdleTimeout:                  0,
		ReadBufferSize:               0,
		WriteBufferSize:              0,
		CompressedFileSuffix:         "",
		ProxyHeader:                  "",
		GETOnly:                      false,
		ErrorHandler:                 nil,
		DisableKeepalive:             false,
		DisableDefaultDate:           false,
		DisableDefaultContentType:    false,
		DisableHeaderNormalizing:     false,
		DisableStartupMessage:        false,
		AppName:                      "NonameYoungService",
		StreamRequestBody:            false,
		DisablePreParseMultipartForm: false,
		ReduceMemoryUsage:            false,
		JSONEncoder:                  nil,
		JSONDecoder:                  nil,
		XMLEncoder:                   nil,
		Network:                      "",
		EnableTrustedProxyCheck:      false,
		TrustedProxies:               nil,
		EnableIPValidation:           false,
		EnablePrintRoutes:            false,
		ColorScheme:                  fiber.Colors{},
		RequestMethods:               nil,
	})
	app.Use(logger.New(logger.Config{
		Next:         nil,
		Done:         nil,
		CustomTags:   nil,
		Format:       "${time} | ${status}  ${latency} | ${method} | ${path}\n",
		TimeFormat:   time.RFC822,
		TimeZone:     "",
		TimeInterval: 0,
		Output:       os.Stdout,
	}))
	handlers.InitHandlers(app)
	return app
}
