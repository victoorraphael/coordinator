package connect

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
)

func MongoConnect() (*mongo.Database, error) {
	connStr := os.Getenv("MONGO_URI")

	if connStr == "" {
		return nil, errors.New("empty mongo db connection string")
	}

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(connStr))
	if err != nil {
		return nil, err
	}

	db := client.Database("schoolplus")
	return db, nil
}
