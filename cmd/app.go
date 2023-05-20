package main

import (
	"InternService/config"
	"InternService/internal/client"
	"InternService/pkg/logger"
	"InternService/pkg/postgresql"
	"context"
)

func main() {
	ctx := context.Background()
	log := logger.GetLogger()
	Config := config.ParseConfig(config.LoadConfig())
	pool := postgresql.GetPool(ctx, Config)
	if err := pool.Ping(ctx); err != nil {
		log.Fatal().Err(err).Msg("unable to get pool")
	}

	postgresql.InitMigrationsOnStart(Config)
	app := client.NewClient(ctx, Config)
	log.Fatal().Err(app.Listen(":8080")).Msg("Something went wrong")
}
