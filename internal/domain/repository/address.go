package repository

import (
	"context"
	"github.com/victoorraphael/coordinator/internal/domain"
)

type IAddress interface {
	Find(ctx context.Context, id int64) (domain.Address, error)
	Add(ctx context.Context, addr *domain.Address) error
}
