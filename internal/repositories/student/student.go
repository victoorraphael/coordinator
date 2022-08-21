package student

import (
	"errors"
	"github.com/victoorraphael/school-plus-BE/infra/entities"

	"go.mongodb.org/mongo-driver/bson/primitive"
	_ "go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	ErrInvalidPerson    = errors.New("a student has to have a valid person with name and email")
	ErrStudentNotFound  = errors.New("student not found")
	ErrFailedAddStudent = errors.New("failed to create student")
	ErrUpdateStudent    = errors.New("failed to  update student")
)

type Repository interface {
	List() ([]Student, error)
	Get(primitive.ObjectID) (Student, error)
	Add(*Student) (map[string]interface{}, error)
	Update(Student) error
	Delete(primitive.ObjectID) error
}

type Student struct {
	ID primitive.ObjectID `json:"id" bson:"_id"`
	entities.Person
}

func New(p entities.Person) (Student, error) {
	if p.Email == "" || p.Name == "" {
		return Student{}, ErrInvalidPerson
	}

	uid := primitive.NewObjectID()

	return Student{
		ID: uid,
		Person: entities.Person{
			ID:      uid,
			Name:    p.Name,
			Email:   p.Email,
			Phone:   p.Phone,
			Address: p.Address,
		},
	}, nil
}
