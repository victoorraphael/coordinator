package mongo

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoRepository struct {
	db         *mongo.Database
	collection *mongo.Collection
}

func New(connStr string) (*MongoRepository, error) {
	if connStr == "" {
		return nil, errors.New("empty mongo db connection string")
	}

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(connStr))
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

func (repo *MongoRepository) Ping() bool {
	log.Println("trying to ping database...")
	err := repo.db.Client().Ping(context.TODO(), readpref.Primary())
	if err != nil {
		log.Println("failed to ping mongo db, err:", err)
		return false
	}
	return true
}
