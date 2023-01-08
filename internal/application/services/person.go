package services

import (
	"context"

	"github.com/victoorraphael/coordinator/internal/domain"
	"github.com/victoorraphael/coordinator/internal/domain/repository"
)

func ListPerson(p repository.IPerson) ([]domain.Person, error) {
	return p.List(context.Background(), domain.Person{})
}
