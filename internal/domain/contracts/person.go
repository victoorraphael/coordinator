package contracts

import (
	"context"
	"github.com/google/uuid"
	"github.com/victoorraphael/coordinator/internal/adapters/postgres/models"
	"github.com/victoorraphael/coordinator/internal/domain"
)

type IPerson interface {
	List(ctx context.Context, person models.Person) ([]domain.Person, error)
	Add(ctx context.Context, person models.Person) (uuid.UUID, error)
}
