package mongo

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoRepository struct {
	db         *mongo.Database
	collection *mongo.Collection
}

func New(ctx context.Context, connStr string) (*MongoRepository, error) {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(connStr))
	if err != nil {
		return nil, err
	}

	db := client.Database("schoolplus")
	collection := db.Collection("student")

	return &MongoRepository{
		db:         db,
		collection: collection,
	}, nil
}
