package repository

import (
	"context"
	"github.com/victoorraphael/coordinator/internal/domain/entities"
	"github.com/victoorraphael/coordinator/pkg/database"
)

type ISubjectRepository interface {
	Add(ctx context.Context, subject *entities.Subject) error
	Delete(ctx context.Context, id int64) error
}

type subject struct {
	pool database.DBPool
}

func NewSubjectRepository(pool database.DBPool) ISubjectRepository {
	return &subject{pool}
}

func (s *subject) Delete(ctx context.Context, id int64) error {
	conn, err := s.pool.Acquire()
	if err != nil {
		return err
	}
	defer s.pool.Release(conn)

	_, err = conn.DeleteFrom("subjects").
		Where("id = ?", id).
		ExecContext(ctx)
	return err
}

func (s *subject) Add(ctx context.Context, subject *entities.Subject) error {
	conn, err := s.pool.Acquire()
	if err != nil {
		return err
	}
	defer s.pool.Release(conn)

	err = conn.InsertInto("subjects").
		Pair("nome", subject.Name).
		Returning("id").
		LoadContext(ctx, &subject.ID)
	return err
}
