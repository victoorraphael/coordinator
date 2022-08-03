package student

import (
	"errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
	_ "go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/victoorraphael/school-plus-BE/domain/entities"
)

var (
	ErrInvalidPerson    = errors.New("a student has to have a valid person with name and email")
	ErrStudentNotFound  = errors.New("student not found")
	ErrFailedAddStudent = errors.New("failed to create student")
	ErrUpdateStudent    = errors.New("failed to  update student")
)

type StudentRepository interface {
	Get(primitive.ObjectID) (Student, error)
	Add(*Student) (map[string]interface{}, error)
	Update(Student) error
	Delete(primitive.ObjectID) error
}

type Student struct {
	person entities.Person
}

func New(p entities.Person) (Student, error) {
	if p.ID == primitive.NilObjectID {
		p.ID = primitive.NewObjectID()
	}

	if p.Email == "" || p.Name == "" {
		return Student{}, ErrInvalidPerson
	}

	return Student{
		person: p,
	}, nil
}

func (s *Student) GetID() primitive.ObjectID {
	return s.person.ID
}

func (s *Student) SetID(id primitive.ObjectID) {
	s.person.ID = id
}

func (s *Student) GetName() string {
	return s.person.Name
}

func (s *Student) SetName(name string) *Student {
	s.person.Name = name
	return s
}

func (s *Student) GetEmail() string {
	return s.person.Email
}

func (s *Student) SetEmail(email string) *Student {
	s.person.Email = email
	return s
}
