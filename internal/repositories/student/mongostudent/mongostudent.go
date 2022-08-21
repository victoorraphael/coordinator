package mongostudent

import (
	"context"
	adapters2 "github.com/victoorraphael/school-plus-BE/infra/adapters"
	"github.com/victoorraphael/school-plus-BE/infra/entities"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"

	"github.com/victoorraphael/school-plus-BE/internal/repositories/student"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository struct {
	collection *mongo.Collection
}

func New(adapters *entities.Adapters) Repository {
	adapter := adapters.DB.(*adapters2.MongoAdapter)
	return Repository{
		collection: adapter.GetCollection("students"),
	}
}

//List retrieve students list from db
func (repo *Repository) List() ([]student.Student, error) {
	students := make([]student.Student, 0)
	cur, err := repo.collection.Find(context.TODO(), bson.D{})
	if err != nil {
		log.Println("error: failed to retrieve students from db: err:", err)
		return students, err
	}

	defer func(cur *mongo.Cursor, ctx context.Context) {
		err := cur.Close(ctx)
		if err != nil {
			log.Fatal("error: failed to close cursor: err:", err)
		}
	}(cur, context.Background())

	for cur.Next(context.Background()) {
		std := student.Student{}
		err := cur.Decode(&std)
		if err != nil {
			log.Println("error: failed to decode data into students struct: err:", err)
			return students, err
		}

		students = append(students, std)
	}

	return students, nil
}

func (repo *Repository) Get(id primitive.ObjectID) (student.Student, error) {
	log.Println("info: looking for student with ID:", id)
	var resp student.Student
	err := repo.collection.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&resp)
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

func (repo *Repository) Add(student *student.Student) (map[string]interface{}, error) {
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

func (repo *Repository) Update(student student.Student) error {
	//TODO implement me
	panic("implement me")
}

func (repo *Repository) Delete(id primitive.ObjectID) error {
	//TODO implement me
	panic("implement me")
}
