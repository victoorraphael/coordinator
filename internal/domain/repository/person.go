package repository

import (
	"context"
	"github.com/victoorraphael/coordinator/internal/domain"
)

type IPerson interface {
	List(ctx context.Context, person domain.Person) ([]domain.Person, error)
	Add(ctx context.Context, person domain.Person) (domain.Person, error)
}
