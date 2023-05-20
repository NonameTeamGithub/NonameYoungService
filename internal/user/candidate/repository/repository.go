package repository

import (
	"database/sql"
)

type CandidateRepository struct {
	//storage.Storage
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) *CandidateRepository {
	return &CandidateRepository{DB: db}
}

//func (r *CandidateRepository) Create(c *fiber.Ctx) error {
//	// реализация метода создания пользователя в базе данных
//	log := logger.GetLogger()
//	var user candidate.Candidate
//	if err := c.BodyParser(&user); err != nil {
//		log.Warn().Err(err)
//		c.Status(400).JSON(&fiber.Map{
//			"success": false,
//			"message": err,
//		})
//		return err
//	}
//	res, err := r.DB.Query("INSERT INTO candidates (id_candidate, name_candidate, surname_candidate, email, password) VALUES ($1, $2, $3, $4, $5)", user.Id_candidate, user.Name_candidate, user.Surname_candidate, user.Email, user.Password)
//	log.Print(res)
//	if err != nil {
//		log.Err(err)
//		c.Status(500).JSON(&fiber.Map{
//			"success": false,
//			"message": err,
//		})
//		return err
//	}
//	// Print result
//
//	// Return Product in JSON format
//	if err := c.JSON(&fiber.Map{
//		"success": true,
//		"message": "User successfully created",
//	}); err != nil {
//		log.Err(err)
//		c.Status(500).JSON(&fiber.Map{
//			"success": false,
//			"message": "Error creating product",
//		})
//		return err
//	}
//	return nil
//}

//	func (r *CandidateRepository) Update(c *fiber.Ctx) error {
//		// реализация метода обновления пользователя в базе данных
//	}
//
//	func (r *CandidateRepository) Delete(id int) error {
//		// реализация метода удаления пользователя из базы данных
//	}
//func (r *CandidateRepository) GetById(c *fiber.Ctx) error {
//	// реализация метода получения пользователя по id из базы данных
//	log := logger.GetLogger()
//	cand := candidate.Candidate{}
//	id := c.Params("id")
//	row := r.DB.QueryRow("SELECT * FROM candidates WHERE id_candidate = $1", id)
//	switch err := row.Scan(&cand.Id_candidate, &cand.Name_candidate, &cand.Surname_candidate, &cand.Email, &cand.Password); err {
//	case sql.ErrNoRows:
//		log.Info().Msg("No rows were returned!")
//		c.Status(500).JSON(&fiber.Map{
//			"success": false,
//			"message": err,
//		})
//	default:
//		//   panic(err)
//		c.Status(500).JSON(&fiber.Map{
//			"success": false,
//			"message": err,
//		})
//	}
//	if err := c.JSON(&fiber.Map{
//		"success": true,
//		"message": "Successfully fetched product",
//		"product": cand,
//	}); err != nil {
//		c.Status(500).JSON(&fiber.Map{
//			"success": false,
//			"message": err,
//		})
//	}
//	return nil
//}

//
//func (r *CandidateRepository) GetAll() ([]*candidate.Candidate, error) {
//	// реализация метода получения всех пользователей из базы данных
//}
