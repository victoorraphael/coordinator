package service

import (
	"context"
	"github.com/victoorraphael/coordinator/internal/entities"
	"log"
)

type IStudentSRV interface {
	Add(ctx context.Context, s entities.Student) (entities.Student, error)
	List(ctx context.Context) ([]entities.Student, error)
	Get(ctx context.Context, s entities.Student) (entities.Student, error)
}

func NewStudentService(adapters *entities.Adapters) IStudentSRV {
	return &student{adapters: adapters}
}

type student struct {
	adapters *entities.Adapters
}

func (srv *student) Add(ctx context.Context, s entities.Student) (entities.Student, error) {
	db := srv.adapters.DB.GetDatabase()
	stmt, err := db.PrepareContext(ctx, ``)
	if err != nil {
		log.Fatal(err)
	}
	if err := stmt.QueryRowContext(ctx).Scan(); err != nil {
		log.Fatal(err)
	}

	return entities.Student{}, err
}

func (srv *student) List(ctx context.Context) ([]entities.Student, error) {
	//TODO implement me
	panic("implement me")
}

func (srv *student) Get(ctx context.Context, s entities.Student) (entities.Student, error) {
	//TODO implement me
	panic("implement me")
}
