package repository

import (
	"InternService/internal/storage"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoInstance struct {
	Client   *mongo.Client
	Database *mongo.Database
}

func NewMongoStorage(instance MongoInstance) storage.MongoStorage {
	return &instance
}
