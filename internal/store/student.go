package store

import (
	"context"
	"github.com/victoorraphael/coordinator/internal/entities"
)

type studentStore struct {
	adapters *entities.Adapters
	person   personStore
}

func (s *studentStore) List(ctx context.Context) ([]entities.Student, error) {
	persons, err := s.person.List(ctx, entities.Person{Type: entities.PersonStudent})
	if err != nil {
		return nil, err
	}

	students := make([]entities.Student, 0)
	for _, p := range persons {
		students = append(students, entities.Student{Person: p})
	}
	return students, nil
}

func (s *studentStore) Add(ctx context.Context, std entities.Student) (entities.Student, error) {
	p, err := s.person.Add(ctx, std.Person)
	if err != nil {
		return entities.Student{}, err
	}

	student := entities.NewStudent()
	student.Person = p
	return student, nil
}

func (s *studentStore) FindByEmail(ctx context.Context, student entities.Student) (entities.Student, error) {
	person, err := s.person.FindByField(ctx, "email", student.Email)
	return entities.Student{Person: person}, err
}

func (s *studentStore) FindByUUID(ctx context.Context, student entities.Student) (entities.Student, error) {
	person, err := s.person.FindByField(ctx, "uuid", student.UUID)
	return entities.Student{Person: person}, err
}
