package contracts

import (
	"github.com/google/uuid"
	"github.com/victoorraphael/coordinator/internal/adapters/postgres/models"
	"github.com/victoorraphael/coordinator/internal/domain"
)

type PersonRepo interface {
	List(person models.Person) ([]domain.Person, error)
	Add(person models.Person) (uuid.UUID, error)
}

type PersonService interface {
}
