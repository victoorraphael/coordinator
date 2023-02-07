package services

import (
	"context"
	"github.com/victoorraphael/coordinator/internal/domain/entities"
)

type IStudentService interface {
	Create(ctx context.Context, std entities.Student) error
	Find(ctx context.Context, std entities.Student) ([]entities.Student, error)
	Update(ctx context.Context, uuid string) error
	Delete(ctx context.Context, uuid string) error
}
