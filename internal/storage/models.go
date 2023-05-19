package storage

import "time"

type Storage interface {
	SelectInternById() (user *User, err error)
	SelectCanditateById()
	SelectCuratorById()
	SelectMentorById()
	SelectHrById()
	InsertInternById()
	InsertCanditateById()
	InsertCuratorById()
	InsertMentorById()
	InsertHrById()
	UpdateInternById()
	UpdateCandidateById()
	UpdateCuratorById()
	UpdateMentorById()
	UpdateHrById()
	DeleteInternById()
	DeleteCanditateById()
	DeleteCuratorById()
	DeleteMentorById()
	DeleteHrById()
}

type User struct {
	Id        int64
	Name      string
	Surname   string
	Email     string
	Password  string
	CreatedAt time.Time
}

type Intern struct {
	User
}
