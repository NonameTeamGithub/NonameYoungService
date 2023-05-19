package mongodb

import (
	"InternService/config"
	"InternService/internal/storage/mongodb/repository"
	"InternService/pkg/logger"
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

func GetMongoConn(ctx context.Context, c config.Config) (repository.MongoInstance, error) {
	log := logger.GetLogger()
	client, clientError := mongo.NewClient(options.Client().ApplyURI(c.MongoDB.DatabaseConnection))
	if clientError != nil {
		log.Warn().Err(clientError).Msg("Unable to get mongoClient.")
		return repository.MongoInstance{}, clientError
	}
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()
	connectionError := client.Connect(ctx)
	db := client.Database(c.MongoDB.DatabaseName)

	if connectionError != nil {
		log.Warn().Err(connectionError).Msg("Unable to connect to mongo db.")
		return repository.MongoInstance{}, connectionError
	}
	log.Info().Msg("Connected to mongodb successfully.")
	return repository.MongoInstance{
		Client:   client,
		Database: db,
	}, nil
}
