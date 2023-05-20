package repository

import (
	"github.com/jackc/pgx/v5/pgxpool"
)

type Storage struct {
	PostgresqlPoolInstance *pgxpool.Pool
}

//
//func (s *Storage) GetById(ctx context.Context) (candidate.Candidate, error) {
//	conn, err := s.PostgresqlPoolInstance.Acquire(ctx)
//	if err != nil {
//
//	}
//	defer conn.Release()
//	q := ``
//	conn.Exec()
//
//}

func NewPostgresqlStorage(psql *pgxpool.Pool) *Storage {
	return &Storage{
		PostgresqlPoolInstance: psql,
	}
}
