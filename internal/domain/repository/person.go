package repository

import (
	"github.com/google/uuid"
	"github.com/victoorraphael/coordinator/internal/domain/entities"
	"github.com/victoorraphael/coordinator/pkg/database"
)

type PersonRepo interface {
	List(person entities.Person) ([]entities.Person, error)
	Add(person entities.Person) (uuid.UUID, error)
}

type person struct {
	pool database.DBPool
}

func NewPersonRepo(pool database.DBPool) PersonRepo {
	return &person{pool}
}

func (p person) List(person entities.Person) ([]entities.Person, error) {
	//TODO implement me
	panic("implement me")
}

func (p person) Add(person entities.Person) (uuid.UUID, error) {
	//TODO implement me
	panic("implement me")
}
