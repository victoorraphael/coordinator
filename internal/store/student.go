package store

import (
	"context"
	"github.com/victoorraphael/coordinator/internal/adapters"
	"github.com/victoorraphael/coordinator/internal/entities"
	"github.com/victoorraphael/coordinator/internal/student"
)

type studentStore struct {
	adapters *adapters.Adapters
	person   personStore
}

func (s *studentStore) List(ctx context.Context) ([]student.Student, error) {
	persons, err := s.person.List(ctx, entities.Person{Type: entities.PersonStudent})
	if err != nil {
		return nil, err
	}

	students := make([]student.Student, 0)
	for _, p := range persons {
		students = append(students, student.Student{Person: p})
	}
	return students, nil
}

func (s *studentStore) Add(ctx context.Context, std student.Student) (student.Student, error) {
	p, err := s.person.Add(ctx, std.Person)
	if err != nil {
		return student.Student{}, err
	}

	student := student.NewStudent()
	student.Person = p
	return student, nil
}

func (s *studentStore) FindByEmail(ctx context.Context, student student.Student) (student.Student, error) {
	person, err := s.person.FindByField(ctx, "email", student.Email)
	return student.Student{Person: person}, err
}

func (s *studentStore) FindByUUID(ctx context.Context, student student.Student) (student.Student, error) {
	person, err := s.person.FindByField(ctx, "uuid", student.UUID)
	return student.Student{Person: person}, err
}

func (s *studentStore) Delete(ctx context.Context, student student.Student) error {
	return s.person.Delete(ctx, student)
}

func (s *studentStore) Update(ctx context.Context, student student.Student) error {
	return s.person.Update(ctx, student)
}
