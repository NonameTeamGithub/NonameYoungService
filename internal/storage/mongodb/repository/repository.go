package repository

import (
	"InternService/internal/auth"
	"InternService/internal/storage"
	"InternService/internal/utilities"
	"InternService/internal/utilities/constants"
	"InternService/internal/utilities/jwtokens"
	"errors"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"os"
	"strconv"
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
	UserCollection := m.Database.Collection("users")
	existingRecord := UserCollection.FindOne(
		ctx.Context(),
		bson.D{{Key: "email", Value: email}},
	)
	existingUser := auth.User{}
	existingRecord.Decode(&existingUser)
	if existingUser.ID != "" {
		return nil
	}
	return errors.New("UserNotFound")
}

func (m MongoInstance) InsertUser(ctx *fiber.Ctx, NewUser auth.User, password string) (string, *auth.User, error) {
	UserCollection := m.Database.Collection("users")
	insertionResult, insertionError := UserCollection.InsertOne(ctx.Context(), NewUser)
	if insertionError != nil {
		return "", &auth.User{}, errors.New("mongo. unable insert user")
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
	if hashError != nil {
		return "", &auth.User{}, errors.New(constants.ResponseMessages.InternalServerError)
	}

	// create a new Password record and insert it
	now := utilities.MakeTimestamp()
	NewPassword := new(auth.Password)
	NewPassword.Created = now
	NewPassword.Hash = hash
	NewPassword.ID = ""
	NewPassword.Updated = now
	NewPassword.UserId = createdUser.ID
	_, insertionError = PasswordCollection.InsertOne(ctx.Context(), NewPassword)
	if insertionError != nil {
		return "", &auth.User{}, errors.New(constants.ResponseMessages.InternalServerError)
	}
	accessExpiration, expirationError := strconv.Atoi(os.Getenv("TOKENS_ACCESS_EXPIRATION"))
	if expirationError != nil {
		accessExpiration = 24
	}
	token, tokenError := jwtokens.GenerateJWT(jwtokens.GenerateJWTParams{
		ExpiresIn: int64(accessExpiration),
		UserId:    createdUser.ID,
	})
	if tokenError != nil {
		return "", &auth.User{}, errors.New(constants.ResponseMessages.InternalServerError)
	}
	return token, createdUser, nil
}

func NewMongoStorage(instance MongoInstance) storage.MongoStorage {
	return &instance
}
