package main

import (
	"InternService/config"
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
	//postgresql.MigratesDown(ctx, Config)
	postgresql.MigratesUp(ctx, Config)
}
