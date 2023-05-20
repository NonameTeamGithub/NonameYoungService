package candidate

import "github.com/gofiber/fiber/v2"

type CandidateI interface {
	Create(ctx *fiber.Ctx) error
}
type Candidate struct {
	Id_candidate      string
	Name_candidate    string
	Surname_candidate string
	Email             string
	Password          string
}
