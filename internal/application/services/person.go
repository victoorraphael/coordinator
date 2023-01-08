package services

import (
	"context"
	"github.com/google/uuid"
	"github.com/victoorraphael/coordinator/internal/adapters/postgres/models"
	"github.com/victoorraphael/coordinator/internal/domain"
	"github.com/victoorraphael/coordinator/internal/domain/repository"
)

func ListPerson(ctx context.Context, p repository.IPerson) ([]domain.Person, error) {
	person := models.Person{Type: domain.PersonStudent}
	return p.List(ctx, person)
}

func CreatePerson(ctx context.Context, p repository.IPerson, person *domain.Student) (uuid.UUID, error) {
	payload := models.Person{
		Name:      person.Name,
		Email:     person.Email,
		Phone:     person.Phone,
		Type:      domain.PersonStudent,
		Birthdate: person.Birthdate,
		AddressID: 1,
	}
	return p.Add(ctx, payload)
}
