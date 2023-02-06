package services

import (
	"context"
	"github.com/victoorraphael/coordinator/internal/domain/entities"
	"github.com/victoorraphael/coordinator/internal/domain/repository"
)

type IPersonService interface {
	FetchAll(ctx context.Context, t entities.PersonType) ([]entities.Person, error)
	Create(ctx context.Context, person entities.Person) (int64, error)
}

func NewPersonService(repo *repository.Repo) IPersonService {
	return &person{repo}
}

type person struct {
	repo *repository.Repo
}

func (p *person) FetchAll(ctx context.Context, t entities.PersonType) ([]entities.Person, error) {
	return p.repo.Person.List(ctx, t)
}

func (p *person) Create(ctx context.Context, person entities.Person) (int64, error) {
	err := p.repo.Person.Add(ctx, &person)
	return person.ID, err
}
