package student

import (
	"errors"

	"github.com/google/uuid"
	"github.com/victoorraphael/school-plus-BE/domain/entities"
)

var (
	ErrInvalidPerson    = errors.New("a student has to have a valid person")
	ErrStudentNotFound  = errors.New("student not found")
	ErrFailedAddStudent = errors.New("failed to create student")
	ErrUpdateStudent    = errors.New("failed to  update student")
)

type StudentRepository interface {
	Get(uuid.UUID) (Student, error)
	Add(Student) error
	Update(Student) error
	Delete(uuid.UUID) error
}

type Student struct {
	person entities.Person
}

func New(p entities.Person) (Student, error) {
	if p.ID == uuid.Nil {
		p.ID = uuid.New()
	}

	if p.Email == "" || p.Name == "" {
		return Student{}, ErrInvalidPerson
	}

	return Student{
		person: p,
	}, nil
}

func (s *Student) GetID() uuid.UUID {
	return s.person.ID
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
