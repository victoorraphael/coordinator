package repository

import (
	"context"
	"github.com/gocraft/dbr/v2"
	"github.com/victoorraphael/coordinator/internal/domain/entities"
	"github.com/victoorraphael/coordinator/pkg/database"
)

type ISchoolRepository interface {
	Add(ctx context.Context, school *entities.School) error
	Delete(ctx context.Context, id int64) error
	FindClassrooms(ctx context.Context, id int64) ([]entities.Classroom, error)
}

func NewSchoolRespository(pool database.DBPool) ISchoolRepository {
	return &school{pool}
}

type school struct {
	pool database.DBPool
}

func (s *school) FindClassrooms(ctx context.Context, id int64) ([]entities.Classroom, error) {
	conn, err := s.pool.Acquire()
	if err != nil {
		return nil, err
	}
	defer s.pool.Release(conn)

	res := make([]entities.Classroom, 0)
	_, err = conn.Select("*").
		From("classroom").
		Join(dbr.I("school_classroom").As("sc"), "sc.classroom_id = id").
		Where("school_id = ?", id).
		LoadContext(ctx, &res)
	return res, err
}

func (s *school) Add(ctx context.Context, school *entities.School) error {
	conn, err := s.pool.Acquire()
	if err != nil {
		return err
	}
	defer s.pool.Release(conn)

	err = conn.InsertInto("school").
		Pair("name", school.Name).
		Pair("address_id", school.Address.ID).
		Returning("id").
		LoadContext(ctx, &school.ID)
	return err
}

func (s *school) Delete(ctx context.Context, id int64) error {
	conn, err := s.pool.Acquire()
	if err != nil {
		return err
	}
	defer s.pool.Release(conn)

	_, err = conn.DeleteFrom("school").
		Where("id = ?", id).
		ExecContext(ctx)
	return err
}
