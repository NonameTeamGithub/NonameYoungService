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
	InsertUser(ctx *fiber.Ctx, NewUser auth.User, password string) (string, *auth.User, error)
}

type PostgresqlStorage interface {
}
