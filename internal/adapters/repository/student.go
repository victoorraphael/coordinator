package repository

import (
	"context"
	"github.com/google/uuid"
	"github.com/victoorraphael/coordinator/internal/domain"
	"time"
)

type Student struct{}

func (s *Student) Find(ctx context.Context, schoolID int, classroomsID ...int) ([]domain.Student, error) {
	return []domain.Student{{
		Classroom: domain.Classroom{},
		School:    domain.School{},
		Person: domain.Person{
			ID:        0,
			UUID:      uuid.UUID{},
			Name:      "teste",
			Email:     "raphael",
			Phone:     "9991291293",
			Birthdate: time.Time{},
			Type:      0,
			Address:   domain.Address{},
		},
	}}, nil
}

func (s *Student) FindUUID(ctx context.Context, uuid string) (domain.Student, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Student) Create(ctx context.Context, student domain.Student) (domain.Student, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Student) Update(ctx context.Context, student domain.Student) error {
	//TODO implement me
	panic("implement me")
}

func (s *Student) Delete(ctx context.Context, uuid string) error {
	//TODO implement me
	panic("implement me")
}
