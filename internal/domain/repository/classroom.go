package repository

import (
	"context"
	"github.com/gocraft/dbr/v2"
	"github.com/victoorraphael/coordinator/internal/domain/entities"
	"github.com/victoorraphael/coordinator/pkg/database"
)

type IClassroomRepository interface {
	Add(ctx context.Context, c *entities.Classroom) error
	Delete(ctx context.Context, id int64) error
	FindSubjects(ctx context.Context, id int64) ([]entities.Subject, error)
}

func NewClassroomRepository(pool database.DBPool) IClassroomRepository {
	return &classroom{pool}
}

type classroom struct {
	pool database.DBPool
}

func (cl *classroom) FindSubjects(ctx context.Context, id int64) ([]entities.Subject, error) {
	conn, err := cl.pool.Acquire()
	if err != nil {
		return nil, err
	}
	defer cl.pool.Release(conn)

	res := make([]entities.Subject, 0)
	_, err = conn.Select("*").
		From("subjects_classroom").
		Join(dbr.I("subjects").As("sub"), "sub.id = subject_id").
		Where("classroom_id = ?", id).
		LoadContext(ctx, &res)
	return res, err
}

func (cl *classroom) Add(ctx context.Context, c *entities.Classroom) error {
	conn, err := cl.pool.Acquire()
	if err != nil {
		return err
	}
	defer cl.pool.Release(conn)

	err = conn.InsertInto("classroom").
		Pair("name", c.Name).
		Returning("id").
		LoadContext(ctx, &c.ID)
	return err
}

func (cl *classroom) Delete(ctx context.Context, id int64) error {
	conn, err := cl.pool.Acquire()
	if err != nil {
		return err
	}
	defer cl.pool.Release(conn)

	_, err = conn.DeleteFrom("classroom").
		Where("id = ?", id).
		ExecContext(ctx)
	return err
}
