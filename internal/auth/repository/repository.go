package repository

import (
	"InternService/internal/auth"
	"InternService/internal/storage"
	"InternService/internal/utilities/constants"
	"InternService/internal/utilities/response"
	"InternService/pkg/logger"
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"
)

type Auth struct {
	logger         zerolog.Logger
	MongoInterface storage.MongoStorage
	Psql           storage.PostgresqlStorage
}

func (a Auth) CreateUser(ctx *fiber.Ctx, user auth.User, password string) (string, *auth.User, error) {
	token, newUser, err := a.MongoInterface.InsertUser(ctx, user, password)
	if err != nil {
		return "", &auth.User{}, response.Response(response.ResponseParams{
			Ctx:    ctx,
			Info:   constants.ResponseMessages.InternalServerError,
			Status: fiber.StatusInternalServerError,
		})
	}
	return token, newUser, nil
}

func (a Auth) GetUserById(ctx context.Context, id int) error {
	//TODO implement me
	panic("implement me")
}

func (a Auth) GetUserByEmail(ctx *fiber.Ctx, email string) error {
	err := a.MongoInterface.SelectUserByEmail(ctx, email)
	if err != nil {
		a.logger.Warn().Err(err)
		return err
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

func NewAuthRepository(mongo storage.MongoStorage) auth.AuthRepository {
	return &Auth{MongoInterface: mongo, logger: logger.GetLogger()}
}
