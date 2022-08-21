package adapters

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"os"
)

type MongoAdapter struct {
	db *mongo.Database
}

func (ma MongoAdapter) Migrate() {}

func (ma MongoAdapter) Ping() bool {
	log.Println("trying to ping database...")
	err := ma.db.Client().Ping(context.TODO(), readpref.Primary())
	if err != nil {
		log.Println("failed to ping mongo db, err:", err)
		return false
	}
	return true
}

func (ma *MongoAdapter) Connect() error {
	connStr := os.Getenv("MONGO_URI")

	if connStr == "" {
		return errors.New("empty mongo db connection string")
	}

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(connStr))
	if err != nil {
		return err
	}

	db := client.Database("schoolplus")
	ma.db = db
	return nil
}

func (ma *MongoAdapter) GetCollection(collection string) *mongo.Collection {
	return ma.db.Collection(collection)
}
