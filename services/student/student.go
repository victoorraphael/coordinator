package student

import (
	"errors"
	"github.com/google/uuid"
	"github.com/victoorraphael/school-plus-BE/internal/repositories/student"
	"log"
)

var (
	ErrMissingID = errors.New("missing ID")
)

type service struct {
	students student.StudentRepository
}

type Service interface {
	Get(uuid.UUID) (student.Student, error)
	Add(student.Student) (map[string]interface{}, error)
}

func New(s student.StudentRepository) Service {
	return &service{students: s}
}

func (s service) Get(id uuid.UUID) (student.Student, error) {
	if id == uuid.Nil {
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
	data, err := s.students.Add(&std)
	if err != nil {
		log.Println("error: student.service.Add, err:", err)
		return nil, err
	}

	log.Println("info: student.service.Add successfully created")
	return data, nil
}
