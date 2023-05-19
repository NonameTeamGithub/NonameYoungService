package repository

import (
	"InternService/internal/auth"
	"InternService/internal/storage"
	"InternService/internal/utilities"
	"errors"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoInstance struct {
	Client   *mongo.Client
	Database *mongo.Database
}

func (m MongoInstance) GetCollection(name string) (collection *mongo.Collection, err error) {
	//TODO implement me
	panic("implement me")
}

func (m MongoInstance) SelectUserById(ctx *fiber.Ctx) {
}

func (m MongoInstance) SelectUserByEmail(ctx *fiber.Ctx, email string) error {
	UserCollection := m.Database.Collection("User")
	existingRecord := UserCollection.FindOne(
		ctx.Context(),
		bson.D{{Key: "email", Value: email}},
	)
	existingUser := &auth.User{}
	existingRecord.Decode(existingUser)
	if existingUser.ID != "" {
		return errors.New("mongo. User not found")
	}
	return nil
}

func (m MongoInstance) InsertUser(ctx *fiber.Ctx, NewUser *auth.User, password string) error {
	UserCollection := m.Database.Collection("User")
	insertionResult, insertionError := UserCollection.InsertOne(ctx.Context(), NewUser)
	if insertionError != nil {
		return errors.New("mongo. unable insert")
	}
	createdRecord := UserCollection.FindOne(
		ctx.Context(),
		bson.D{{Key: "_id", Value: insertionResult.InsertedID}},
	)
	createdUser := &auth.User{}
	createdRecord.Decode(createdUser)
	PasswordCollection := m.Database.Collection("Password")
	// create password hash
	hash, hashError := utilities.MakeHash(password)

	// create a new Password record and insert it
	now := utilities.MakeTimestamp()
	NewPassword := new(auth.Password)
	NewPassword.Created = now
	NewPassword.Hash = hash
	NewPassword.ID = ""
	NewPassword.Updated = now
	NewPassword.UserId = createdUser.ID
}

func NewMongoStorage(instance MongoInstance) storage.MongoStorage {
	return &instance
}
