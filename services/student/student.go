package student

import (
	"errors"
	"github.com/victoorraphael/school-plus-BE/domain/entities"
	"github.com/victoorraphael/school-plus-BE/internal/repositories/student"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
)

var (
	ErrMissingID = errors.New("missing ID")
)

type service struct {
	students student.StudentRepository
}

type Service interface {
	Get(id primitive.ObjectID) (student.Student, error)
	Add(student.Student) (map[string]interface{}, error)
}

func New(s student.StudentRepository) Service {
	return &service{students: s}
}

func (s service) Get(id primitive.ObjectID) (student.Student, error) {
	if id == primitive.NilObjectID {
		log.Println("error: student.service.Get, err:", ErrMissingID)
		return student.Student{}, ErrMissingID
	}

	studentFetched, err := s.students.Get(id)
	if err != nil {
		log.Println("error: student.service.Get, err:", err)
		return student.Student{}, err
	}

	return studentFetched, nil
}

func (s service) Add(std student.Student) (map[string]interface{}, error) {
	stud, err := student.New(entities.Person{
		Name:  std.Name,
		Email: std.Email,
		Phone: std.Phone,
	})
	if err != nil {
		log.Println("error: student.service.Add, err:", err)
		return nil, err
	}
	data, err := s.students.Add(&stud)
	if err != nil {
		log.Println("error: student.service.Add, err:", err)
		return nil, err
	}

	log.Println("info: student.service.Add successfully created")
	return data, nil
}
