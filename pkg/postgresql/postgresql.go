package postgresql

import (
	"InternService/config"
	"InternService/pkg/logger"
	"InternService/pkg/middleware"
	"context"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/jackc/pgx/v5/stdlib"
	"time"
)

func GetPool(ctx context.Context, c config.Config) (connect *pgxpool.Pool) {
	log := logger.GetLogger()
	err := middleware.ConnectWithTries(func() error {
		ctx, cancel := context.WithCancel(ctx)
		defer cancel()
		var err error
		connect, err = pgxpool.New(ctx, fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=%s",
			c.PostgreSQLDB.User,
			c.PostgreSQLDB.Pass,
			c.PostgreSQLDB.Host,
			c.PostgreSQLDB.Port,
			c.PostgreSQLDB.Dbname,
			c.PostgreSQLDB.SSLMode))
		return err
	}, 3, time.Second*3)
	if err != nil {
		log.Fatal().Err(err).Msg("unable to connect to database")
	}
	log.Info().Msg("connected to database successfully")
	return connect
}

func MigratesUp(ctx context.Context, c config.Config) {
	log := logger.GetLogger()
	dbUrl := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=%s",
		c.PostgreSQLDB.User,
		c.PostgreSQLDB.Pass,
		c.PostgreSQLDB.Host,
		c.PostgreSQLDB.Port,
		c.PostgreSQLDB.Dbname,
		c.PostgreSQLDB.SSLMode)
	m, err := migrate.New("file://../migrations/", dbUrl)
	if err != nil {
		log.Fatal().Err(err).Msg("Unable to get migrator.")
	}
	if err := m.Up(); err != nil {
		log.Warn().Err(err).Msg("Unable to up migrations.")
	}
	log.Info().Msg("Migration up done successfully!")
}

func MigratesDown(ctx context.Context, c config.Config) {
	log := logger.GetLogger()
	dbUrl := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=%s",
		c.PostgreSQLDB.User,
		c.PostgreSQLDB.Pass,
		c.PostgreSQLDB.Host,
		c.PostgreSQLDB.Port,
		c.PostgreSQLDB.Dbname,
		c.PostgreSQLDB.SSLMode)
	m, err := migrate.New("file://../migrations/", dbUrl)
	if err != nil {
		log.Fatal().Err(err).Msg("Unable to get migrator.")
	}
	if err := m.Down(); err != nil {
		log.Warn().Err(err).Msg("Unable to down migrations.")
	}
	log.Info().Msg("Migration down successfully!")
}
