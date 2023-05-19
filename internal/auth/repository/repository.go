package repository

import (
	"InternService/internal/auth"
	"InternService/internal/storage"
	"InternService/internal/utilities/constants"
	"InternService/internal/utilities/response"
	"context"
	"github.com/gofiber/fiber/v2"
)

type Auth struct {
	Mongo storage.MongoStorage
	Psql  storage.PostgresqlStorage
}

func (a Auth) CreateUser(ctx *fiber.Ctx, user auth.User) error {
	err := a.Mongo.InsertUser(ctx, user)

	if err != nil {
		return response.Response(response.ResponseParams{
			Ctx:    ctx,
			Info:   constants.ResponseMessages.InternalServerError,
			Status: fiber.StatusInternalServerError,
		})
	}
	return nil
}

func (a Auth) GetUserById(ctx context.Context, id int) error {
	//TODO implement me
	panic("implement me")
}

func (a Auth) GetUserByEmail(ctx *fiber.Ctx, email string) error {
	err := a.Mongo.SelectUserByEmail(ctx, email)
	if err != nil {
		return response.Response(response.ResponseParams{
			Ctx:    ctx,
			Info:   constants.ResponseMessages.EmailAlreadyInUse,
			Status: fiber.StatusBadRequest,
		})
	}
	return nil
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
