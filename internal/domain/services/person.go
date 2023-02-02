package services

import (
	"github.com/victoorraphael/coordinator/internal/domain/entities"
	"github.com/victoorraphael/coordinator/internal/domain/repository"
)

type IPersonService interface {
	FetchAll(t entities.PersonType) ([]entities.Person, error)
	Create(person entities.Person) (string, error)
}

type person struct {
	repo *repository.Repo
}

func (p *person) FetchAll(t entities.PersonType) ([]entities.Person, error) {
	return p.repo.Person.List(t)
}

func (p *person) Create(person entities.Person) (string, error) {
	err := p.repo.Person.Add(&person)
	return person.UUID, err
}

func NewPersonService(repo *repository.Repo) IPersonService {
	return &person{repo}
}
