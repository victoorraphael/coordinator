package student

import (
	"context"
	"github.com/victoorraphael/coordinator/internal/adapters"
	"github.com/victoorraphael/coordinator/internal/address"
	"github.com/victoorraphael/coordinator/internal/person"
)

type store struct {
	adapters *adapters.Adapters
	persons  person.Store
}

func (s *store) Find(ctx context.Context, schoolID int, classroomsID ...int) ([]Student, error) {
	return nil, nil
}

func (s *store) FindUUID(ctx context.Context, uuid string) (Student, error) {
	p, err := s.persons.FindByField(ctx, "uuid", uuid)
	if err != nil {
		return Student{}, err
	}

	std := NewStudent(address.Address{})
	std.Person = p
	return *std, nil
}

func (s *store) Create(ctx context.Context, student Student) (Student, error) {
	p, err := s.persons.Add(ctx, student.Person)
	return Student{Person: p}, err
}

func (s *store) Update(ctx context.Context, student Student) error {
	//TODO implement me
	panic("implement me")
}

func (s *store) Delete(ctx context.Context, uuid string) error {
	//TODO implement me
	panic("implement me")
}

func NewStore(adapters *adapters.Adapters) Repository {
	return &store{
		adapters: adapters,
		persons:  person.Store{Adapters: adapters},
	}
}
