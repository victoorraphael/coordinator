package service

import (
	"context"
	"errors"
	"github.com/victoorraphael/coordinator/internal/entities"
	"github.com/victoorraphael/coordinator/internal/store"
	"time"
)

type IStudentSRV interface {
	Add(ctx context.Context, s entities.Student) (entities.Student, error)
	List(ctx context.Context) ([]entities.Student, error)
	Get(ctx context.Context, s entities.Student) (entities.Student, error)
}

func NewStudentService(store *store.Store) IStudentSRV {
	return &student{store: store}
}

type student struct {
	store *store.Store
}

func (srv *student) Add(ctx context.Context, s entities.Student) (entities.Student, error) {
	res := entities.Student{}
	if s.Name == "" {
		return res, errors.New("name cannot be empty")
	}

	if s.Email == "" {
		return res, errors.New("email cannot be empty")
	}

	if s.Type != entities.PersonStudent {
		return res, errors.New("wrong type")
	}

	if s.Birthdate == "" {
		return res, errors.New("birthdate cannot be empty")
	}

	_, err := time.Parse("2006-01-02", s.Birthdate)
	if err != nil {
		return res, err
	}

	std, err := srv.store.Student.Add(ctx, s)
	if err != nil {
		return entities.Student{}, err
	}
	return std, nil
}

func (srv *student) List(ctx context.Context) ([]entities.Student, error) {
	return srv.store.Student.List(ctx)
}

func (srv *student) Get(ctx context.Context, s entities.Student) (entities.Student, error) {
	//TODO implement me
	panic("implement me")
}
