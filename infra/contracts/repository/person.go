package repository

import (
	"github.com/google/uuid"
	"github.com/victoorraphael/school-plus-BE/infra/entities"
)

type PersonRepo struct {
	DB *entities.Adapters
}

func (p PersonRepo) List() ([]entities.Person, error) {
	panic("implement me")
}

func (p PersonRepo) FindOne(person entities.Person) (entities.Person, error) {
	panic("implement me")
}

func (p PersonRepo) Add(person *entities.Person) (uuid.UUID, error) {
	panic("implement me")
}

func (p PersonRepo) Update(person *entities.Person) error {
	panic("implement me")
}

func (p PersonRepo) Delete(person entities.Person) error {
	panic("implement me")
}
