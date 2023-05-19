package storage

import (
	"InternService/internal/auth"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoStorage interface {
	GetCollection(name string) (collection *mongo.Collection, err error)
	SelectUserById(*fiber.Ctx)
	SelectUserByEmail(ctx *fiber.Ctx, email string) error
	InsertUser(ctx *fiber.Ctx, user auth.User) error
}

type PostgresqlStorage interface {
}
