package models

import (
	"InternService/internal/user/candidate"
	"InternService/pkg/logger"
	"database/sql"
)

type CandidateModel struct {
	DB *sql.DB
	//PostgresqlPoolInstance *pgxpool.Pool
}

func (m *CandidateModel) Create(user candidate.Candidate) error {
	// реализация метода создания пользователя в базе данных
	log := logger.GetLogger()
	stmt := "INSERT INTO candidates (id_candidate, name_candidate, surname_candidate, email, password) VALUES ($1, $2, $3, $4, $5)"
	_, err := m.DB.Query(stmt, user.Id_candidate, user.Name_candidate, user.Surname_candidate, user.Email, user.Password)
	//res, err := m.PostgresqlPoolInstance.Query(c, "INSERT INTO candidates (id_candidate, name_candidate, surname_candidate, email, password) VALUES ($1, $2, $3, $4, $5)", user.Id_candidate, user.Name_candidate, user.Surname_candidate, user.Email, user.Password)
	if err != nil {
		log.Warn().Err(err)
		return err
	}
	return nil
}

func (m *CandidateModel) Update(user candidate.Candidate) error {
	// реализация метода обновления пользователя в базе данных

	return nil
}

func (r *CandidateModel) Delete(id string) error {
	// реализация метода удаления пользователя из базы данных
	stmt := "DELETE FROM candidates WHERE id_candidate = $1"
	_, err := r.DB.Query(stmt, id)
	if err != nil {
		return err
	}
	return nil
}
func (m *CandidateModel) GetById(user *candidate.Candidate, id string) error {
	// реализация метода получения пользователя по id из базы данных
	log := logger.GetLogger()
	stmt := "SELECT * FROM candidates WHERE id_candidate = $1"
	row := m.DB.QueryRow(stmt, id)
	switch err := row.Scan(&user.Id_candidate, &user.Name_candidate, &user.Surname_candidate, &user.Email, &user.Password); err {
	case sql.ErrNoRows:
		log.Warn().Msg("No rows were returned!")
		return err
	default:
		log.Warn().Err(err)
		return err
	}
	return nil
}
