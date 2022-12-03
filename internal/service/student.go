package service

import (
	"context"
	"errors"
	"github.com/victoorraphael/coordinator/internal/entities"
	"github.com/victoorraphael/coordinator/internal/store"
)

type IStudentSRV interface {
	Add(ctx context.Context, s entities.Student) (entities.Student, error)
	List(ctx context.Context) ([]entities.Student, error)
	Get(ctx context.Context, s entities.Student) (entities.Student, error)
	Delete(ctx context.Context, s entities.Student) error
}

func NewStudentService(store *store.Store) IStudentSRV {
	return &student{store: store}
}

type student struct {
	store *store.Store
}

func (srv *student) Delete(ctx context.Context, s entities.Student) error {
	_, err := srv.store.Student.FindByUUID(ctx, s)
	if err != nil {
		return err
	}

	return srv.store.Student.Delete(ctx, s)
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

	if s.Birthdate.IsZero() {
		return res, errors.New("birthdate cannot be empty")
	}

	student, err := srv.store.Student.FindByEmail(ctx, s)
	if err != nil {
		return res, err
	}

	if student.UUID.String() != "" {
		return res, errors.New("student already exists")
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
	return srv.store.Student.FindByUUID(ctx, s)
}
