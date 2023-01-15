package contracts

import (
	"context"
	"github.com/victoorraphael/coordinator/internal/domain"
)

type IStudent interface {
	Find(ctx context.Context, schoolID int, classroomsID ...int) ([]domain.Student, error)
	FindUUID(ctx context.Context, uuid string) (domain.Student, error)
	Create(ctx context.Context, student domain.Student) (domain.Student, error)
	Update(ctx context.Context, student domain.Student) error
	Delete(ctx context.Context, uuid string) error
}
