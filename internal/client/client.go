package client

import (
	"InternService/config"
	"InternService/internal/auth/repository"
	"InternService/internal/auth/usecase"
	"InternService/internal/client/handlers"
	"InternService/pkg/logger"
	"InternService/pkg/mongodb"
	"context"
	"github.com/gofiber/fiber/v2"
	fiberLogger "github.com/gofiber/fiber/v2/middleware/logger"
	"os"
	"time"
)

func NewClient(ctx context.Context, config config.Config) *fiber.App {
	log := logger.GetLogger()
	app := fiber.New(fiber.Config{
		DisableStartupMessage: true,
		AppName:               "NonameYoungService",
	})
	app.Use(fiberLogger.New(fiberLogger.Config{
		Format:     "${time} | ${status}  ${latency} | ${method} | ${path}\n",
		TimeFormat: time.RFC822,
		Output:     os.Stdout,
	}))
	mongo, err := mongodb.GetMongoConn(ctx, config)
	if err != nil {
		log.Warn().Err(err).Msg("Unable to connect to mongo")
	}
	authR := repository.NewAuthRepository(mongo)
	authU := usecase.NewUseCase(authR)
	appContext := handlers.AppContext{log, app, authU}
	handlers.InitHandlers(&appContext)
	return app
}
