package repository

import (
	"InternService/internal/user/candidate"
	"InternService/pkg/logger"
	"database/sql"
	"github.com/gofiber/fiber/v2"
)

type CandidateRepository struct {
	//storage.Storage
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) *CandidateRepository {
	return &CandidateRepository{DB: db}
}

func (r *CandidateRepository) Create(c *fiber.Ctx) error {
	// реализация метода создания пользователя в базе данных
	log := logger.GetLogger()
	var user candidate.Candidate
	if err := c.BodyParser(&user); err != nil {
		log.Warn().Err(err)
		c.Status(400).JSON(&fiber.Map{
			"success": false,
			"message": err,
		})
		return err
	}
	_, err := r.DB.Query("INSERT INTO candidates (id_candidate, name_candidate, surname_candidate, email, password) VALUES ($1, $2, $3, $4)", user.Id_candidate, user.Name_candidate, user.Surname_candidate, user.Email, user.Password)
	if err != nil {
		log.Err(err)
		c.Status(500).JSON(&fiber.Map{
			"success": false,
			"message": err,
		})
		return err
	}
	// Print result

	// Return Product in JSON format
	if err := c.JSON(&fiber.Map{
		"success": true,
		"message": "User successfully created",
	}); err != nil {
		log.Err(err)
		c.Status(500).JSON(&fiber.Map{
			"success": false,
			"message": "Error creating product",
		})
		return err
	}
	return nil
}

//func (r *CandidateRepository) Update(c *fiber.Ctx) error {
//	// реализация метода обновления пользователя в базе данных
//}
//
//func (r *CandidateRepository) Delete(id int) error {
//	// реализация метода удаления пользователя из базы данных
//}
//
//func (r *CandidateRepository) GetById(id int) (*candidate.Candidate, error) {
//	// реализация метода получения пользователя по id из базы данных
//}
//
//func (r *CandidateRepository) GetAll() ([]*candidate.Candidate, error) {
//	// реализация метода получения всех пользователей из базы данных
//}
