package services

import (
	"context"
	"errors"
	"github.com/gocraft/dbr/v2"
	"github.com/golangsugar/chatty"
	"github.com/victoorraphael/coordinator/internal/domain/entities"
	"github.com/victoorraphael/coordinator/internal/domain/repository"
	"github.com/victoorraphael/coordinator/pkg/errs"
	"github.com/victoorraphael/coordinator/pkg/uid"
)

type ISubjectService interface {
	Create(ctx context.Context, sub *entities.Subject) error
	Find(ctx context.Context, sub entities.Subject) (entities.Subject, error)
}

type subject struct {
	repo *repository.Repo
}

func NewSubjectService(repo *repository.Repo) ISubjectService {
	return &subject{repo}
}

func (s *subject) Find(ctx context.Context, sub entities.Subject) (entities.Subject, error) {
	exist, err := s.repo.Subject.Search(ctx, sub)
	if err != nil {
		if errors.Is(err, dbr.ErrNotFound) {
			chatty.Info("subject nao existe")
			return exist, errs.WrapError(errs.ErrFieldViolation, "subject nao existe")
		}
		chatty.Errorf("erro ao buscar subject : err: %v", err)
		return exist, errs.WrapError(errs.ErrInternalError, "nao foi possivel buscar subject")
	}

	return exist, nil
}

func (s *subject) Create(ctx context.Context, sub *entities.Subject) error {
	exist, err := s.repo.Subject.Search(ctx, entities.Subject{
		Name: sub.Name,
	})
	if err != nil && !errors.Is(err, dbr.ErrNotFound) {
		chatty.Errorf("erro ao buscar subject %#v: err: %v", sub, err)
		return errs.WrapError(errs.ErrInternalError, "nao foi possivel criar subject")
	}

	if exist.UUID != "" {
		chatty.Info("subject ja existe")
		return errs.WrapError(errs.ErrFieldViolation, "subject ja existe")
	}

	sub.UUID = uid.NewUUID().String()
	if err := s.repo.Subject.Add(ctx, sub); err != nil {
		chatty.Errorf("erro ao criar subject %#v: err: %v", sub, err)
		return errs.WrapError(errs.ErrInternalError, "nao foi possivel criar subject")
	}
	return nil
}
