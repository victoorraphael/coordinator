package mongo

import (
	"context"
	"errors"
	"log"

	"github.com/google/uuid"
	"github.com/victoorraphael/school-plus-BE/internal/repositories/student"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/readpref"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoRepository struct {
	db         *mongo.Database
	collection *mongo.Collection
}

func (repo *MongoRepository) Get(id uuid.UUID) (student.Student, error) {
	log.Println("info: looking for student with ID:", id)
	var resp student.Student
	err := repo.collection.FindOne(context.TODO(), bson.D{{Key: "_id", Value: id}}).Decode(&resp)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			log.Println("info: doest not exist student with id:", id)
			return resp, err
		}
		log.Println("error: student.Get err:", err)
		return student.Student{}, err
	}

	log.Println("info: student found !")
	return resp, nil
}

func (repo *MongoRepository) Add(student *student.Student) (map[string]interface{}, error) {
	result, err := repo.collection.InsertOne(context.TODO(), student)
	if err != nil {
		log.Println("error: cannot create student, err:", err)
		return nil, err
	}
	log.Println("info: student created with id:", result.InsertedID)
	res := make(map[string]interface{}, 0)
	res["id"] = result.InsertedID
	return res, nil
}

func (repo *MongoRepository) Update(student student.Student) error {
	//TODO implement me
	panic("implement me")
}

func (repo *MongoRepository) Delete(uuid uuid.UUID) error {
	//TODO implement me
	panic("implement me")
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
