package repository

import (
	"InternService/pkg/mongodb"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Storage struct {
	MongoDBInstance        *mongodb.MongoInstance
	PostgresqlPoolInstance *pgxpool.Pool
}

func NewStorage(mdb *mongodb.MongoInstance, psql *pgxpool.Pool) *Storage {
	return &Storage{
		MongoDBInstance:        mdb,
		PostgresqlPoolInstance: psql,
	}
}
