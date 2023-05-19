package repository

import (
	"InternService/internal/auth"
	"InternService/internal/storage"
	"context"
)

type Auth struct {
	Mongo storage.MongoStorage
	Psql  storage.PostgresqlStorage
}

func (a Auth) CreateUser(ctx context.Context, user *auth.User) error {
	panic("implement me")
}

func (a Auth) GetUserById(ctx context.Context, id int) error {
	//TODO implement me
	panic("implement me")
}

func (a Auth) GetUserByEmail(ctx context.Context, email string) error {
	//TODO implement me
	panic("implement me")
}

func (a Auth) Update(ctx context.Context, user *auth.User) error {
	//TODO implement me
	panic("implement me")
}

func (a Auth) Delete(ctx context.Context, user *auth.User) error {
	//TODO implement me
	panic("implement me")
}

func NewAuthRepository(st storage.Storage) auth.AuthRepository {
	return &Auth{st: st}
}
