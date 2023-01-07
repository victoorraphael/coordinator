package student

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/victoorraphael/coordinator/internal/adapters"
	"github.com/victoorraphael/coordinator/internal/address"
	"log"
)

type service struct {
	store   Repository
	address address.Store
}

var (
	ErrInvalidField  = errors.New("campo inválido")
	ErrInternalError = errors.New("internal error")
)

func (s *service) List(ctx context.Context, schoolID int, classroomID int) ([]Student, error) {
	//TODO implement me
	panic("implement me")
}

func (s *service) Get(ctx context.Context, uuid string) (Student, error) {
	//TODO implement me
	panic("implement me")
}

func (s *service) Create(ctx context.Context, student Student) (Student, error) {
	errLog := errors.New("student.Create")
	if student.Email == "" {
		log.Printf("%v: %v: email empty\n", errLog, ErrInvalidField)
		return Student{}, fmt.Errorf("%w: email não pode ser vazio", ErrInvalidField)
	}
	if student.Name == "" {
		log.Printf("%v: %v: nome empty\n", errLog, ErrInvalidField)
		return Student{}, fmt.Errorf("%w: nome não pode ser vazio", ErrInvalidField)
	}
	if student.Birthdate.IsZero() {
		log.Printf("%v: %v: birthdate empty\n", errLog, ErrInvalidField)
		return Student{}, fmt.Errorf("%w: data de nascimento não pode ser vazio", ErrInvalidField)
	}

	addr, err := s.address.Find(ctx, "zip", student.Address.Zip)
	if err != nil && err != sql.ErrNoRows {
		log.Printf("%v: error: %v\n", errLog, err)
		return Student{}, ErrInternalError
	}

	if addr.ID == 0 {
		err := s.address.Create(ctx, &addr)
		if err != nil {
			log.Printf("%v: address.Create: error: %v", errLog, err)
			return Student{}, ErrInternalError
		}
	}

	student.Address = addr
	std, err := s.store.Create(ctx, student)
	if err != nil {
		log.Printf("%v: error: %v\n", errLog, err)
		return Student{}, err
	}

	log.Printf("%v: student created successfully\n", errLog)
	return std, nil
}

func (s *service) Update(ctx context.Context, student Student) error {
	//TODO implement me
	panic("implement me")
}

func (s *service) Delete(ctx context.Context, uuid string) error {
	//TODO implement me
	panic("implement me")
}

func NewService(adapters *adapters.Adapters) Service {
	return &service{
		store:   NewStore(adapters),
		address: address.Store{Adapters: adapters},
	}
}
