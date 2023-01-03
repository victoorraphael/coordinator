package service

import (
	"context"
	"database/sql"
	"errors"
	"github.com/google/uuid"
	"github.com/victoorraphael/coordinator/internal/entities"
	"github.com/victoorraphael/coordinator/internal/store"
	student2 "github.com/victoorraphael/coordinator/internal/student"
)

type IStudentSRV interface {
	Add(ctx context.Context, s student2.Student) (student2.Student, error)
	List(ctx context.Context) ([]student2.Student, error)
	Get(ctx context.Context, s student2.Student) (student2.Student, error)
	Delete(ctx context.Context, s student2.Student) error
	Update(ctx context.Context, s student2.Student) error
}

func NewStudentService(store *store.Store) IStudentSRV {
	return &student{store: store}
}

type student struct {
	store *store.Store
}

func (srv *student) Update(ctx context.Context, s student2.Student) error {
	_, err := srv.store.Student.FindByUUID(ctx, s)
	if err != nil {
		return err
	}

	return srv.store.Student.Update(ctx, s)
}

func (srv *student) Delete(ctx context.Context, s student2.Student) error {
	_, err := srv.store.Student.FindByUUID(ctx, s)
	if err != nil {
		return err
	}

	return srv.store.Student.Delete(ctx, s)
}

func (srv *student) Add(ctx context.Context, s student2.Student) (student2.Student, error) {
	res := student2.Student{}
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
	if err != nil && err != sql.ErrNoRows {
		return res, err
	}

	if student.UUID != uuid.Nil {
		return res, errors.New("student already exists")
	}

	std, err := srv.store.Student.Add(ctx, s)
	if err != nil {
		return student2.Student{}, err
	}
	return std, nil
}

func (srv *student) List(ctx context.Context) ([]student2.Student, error) {
	return srv.store.Student.List(ctx)
}

func (srv *student) Get(ctx context.Context, s student2.Student) (student2.Student, error) {
	return srv.store.Student.FindByUUID(ctx, s)
}
