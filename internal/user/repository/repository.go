package repository

type Repository interface {
	Create(interface{}) error
	Update(interface{}) error
	Delete(interface{}) error
	GetById() (interface{}, error)
	GetAll() (interface{}, error)
}
