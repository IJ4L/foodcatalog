package database

import (
	"context"
	"fmt"

	"github.com/ij4l/foodCatalog/util"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectMongoDb(config util.Config, ctx context.Context) (*mongo.Database, error) {
	client, err := CreateMongoDBConnection(config, ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to create MongoDB client: %v", err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		client.Disconnect(ctx)
		return nil, fmt.Errorf("failed to ping MongoDB: %v", err)
	}

	return client.Database(config.MongoDB), nil
}

func CreateMongoDBConnection(config util.Config, ctx context.Context) (*mongo.Client, error) {
	// credentials := options.Credential{
	// 	Username: config.MongoDBUser,
	// 	Password: config.MongoDBPassword,
	// }

	dsn := fmt.Sprintf("mongodb://%s:%s/", config.MongoDBHost, config.MongoDBPort)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(dsn))
	if err != nil {
		return nil, fmt.Errorf("failed to connect to MongoDB: %v", err)
	}

	return client, nil
}