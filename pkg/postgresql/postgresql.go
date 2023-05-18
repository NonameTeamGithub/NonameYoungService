package postgresql

import (
	"InternService/config"
	"InternService/pkg/logger"
	"InternService/pkg/middleware"
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"time"
)

func GetPool(ctx context.Context, c config.Config) (connect *pgxpool.Pool) {
	log := logger.GetLogger()
	err := middleware.ConnectWithTries(func() error {
		ctx, cancel := context.WithCancel(ctx)
		defer cancel()
		var err error
		connect, err = pgxpool.Connect(ctx, fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?%s&%s",
			c.PostgreSQLDB.User,
			c.PostgreSQLDB.Pass,
			c.PostgreSQLDB.Host,
			c.PostgreSQLDB.Port,
			c.PostgreSQLDB.Dbname,
			c.PostgreSQLDB.SSLMode,
			c.PostgreSQLDB.MaxConns))
		return err
	}, 3, time.Second*3)
	if err != nil {
		log.Fatal().Err(err).Msg("unable to connect to database")
	}
	log.Info().Msg("connected to database successfully")
	return connect
}
