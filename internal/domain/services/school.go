package services

import (
	"context"

	"github.com/victoorraphael/coordinator/internal/domain/entities"
	"github.com/victoorraphael/coordinator/internal/domain/repository"
)

type ISchoolService interface {
	Create(ctx context.Context, item entities.School) (int64, error)
	Search(ctx context.Context, school entities.School) (entities.School, error)
}

type school struct {
	repo *repository.Repo
}

func (s *school) Search(ctx context.Context, school entities.School) (entities.School, error) {
	if school.UUID != "" {
		return s.repo.School.FindUUID(ctx, school.UUID)
	}
	return entities.School{}, nil
}

func NewSchoolService(repo *repository.Repo) ISchoolService {
	return &school{repo}
}

func (s school) Create(ctx context.Context, item entities.School) (int64, error) {
	err := s.repo.School.Add(ctx, &item)
	return item.ID, err
}
