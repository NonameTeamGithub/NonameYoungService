package storage

import "go.mongodb.org/mongo-driver/mongo"

type MongoStorage interface {
	GetCollection(name string) (collection *mongo.Collection, err error)
}

type PostgresqlStorage interface {
}
