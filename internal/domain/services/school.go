package services

import (
	"context"
	"github.com/victoorraphael/coordinator/internal/domain/entities"
	"github.com/victoorraphael/coordinator/internal/domain/repository"
)

type ISchoolService interface {
	Create(ctx context.Context, item entities.School) (int64, error)
}

type school struct {
	repo *repository.Repo
}

func NewSchoolService(repo *repository.Repo) ISchoolService {
	return &school{repo}
}

func (s school) Create(ctx context.Context, item entities.School) (int64, error) {
	err := s.repo.School.Add(ctx, &item)
	return item.ID, err
}
