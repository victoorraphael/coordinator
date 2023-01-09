package repository

import (
	"context"
	"github.com/victoorraphael/coordinator/internal/domain"
)

type ISchool interface {
	Add(ctx context.Context, school domain.School) (int64, error)
}
