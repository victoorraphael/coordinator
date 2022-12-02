package service

import (
	"context"
	"github.com/victoorraphael/coordinator/internal/entities"
	"github.com/victoorraphael/coordinator/internal/store"
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
	return entities.Student{}, nil
}

func (srv *student) List(ctx context.Context) ([]entities.Student, error) {
	return srv.store.Student.List(ctx)
}

func (srv *student) Get(ctx context.Context, s entities.Student) (entities.Student, error) {
	//TODO implement me
	panic("implement me")
}
