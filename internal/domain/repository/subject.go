package repository

import (
	"context"
	"github.com/victoorraphael/coordinator/internal/domain/entities"
	"github.com/victoorraphael/coordinator/pkg/database"
	"github.com/victoorraphael/coordinator/pkg/utils"
)

type ISubjectRepository interface {
	Add(ctx context.Context, subject *entities.Subject) error
	Delete(ctx context.Context, uuid string) error
	Search(ctx context.Context, filter entities.Subject) (entities.Subject, error)
}

type subject struct {
	pool database.DBPool
}

func NewSubjectRepository(pool database.DBPool) ISubjectRepository {
	return &subject{pool}
}

func (s *subject) Search(ctx context.Context, filter entities.Subject) (entities.Subject, error) {
	conn, err := s.pool.Acquire()
	if err != nil {
		return entities.Subject{}, err
	}
	defer s.pool.Release(conn)

	query, values, err := utils.BuildSearchQuery(filter)
	if err != nil {
		return entities.Subject{}, err
	}

	var resp entities.Subject
	_, err = conn.Select("*").
		From("subjects").
		Where(query, values...).
		LoadContext(ctx, &resp)
	return resp, err
}

func (s *subject) Delete(ctx context.Context, uuid string) error {
	conn, err := s.pool.Acquire()
	if err != nil {
		return err
	}
	defer s.pool.Release(conn)

	_, err = conn.DeleteFrom("subjects").
		Where("uuid = ?", uuid).
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
		Pair("uuid", subject.UUID).
		Pair("nome", subject.Name).
		Returning("id").
		LoadContext(ctx, &subject.ID)
	return err
}
