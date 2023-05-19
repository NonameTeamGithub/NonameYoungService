package repository

import (
	"github.com/jackc/pgx/v5/pgxpool"
)

type Storage struct {
	PostgresqlPoolInstance *pgxpool.Pool
}

func NewPostgresqlStorage(psql *pgxpool.Pool) *Storage {
	return &Storage{
		PostgresqlPoolInstance: psql,
	}
}
