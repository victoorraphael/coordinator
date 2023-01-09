package repository

import (
	"context"
	"github.com/victoorraphael/coordinator/internal/adapters/postgres"
	"github.com/victoorraphael/coordinator/internal/domain"
)

type School struct{}

func (s *School) Add(ctx context.Context, school domain.School) (int64, error) {
	db := postgres.NewPostgresAdapter().GetDatabase()
	query := "INSERT INTO school(name) VALUES ($1) RETURNING id"
	result, err := db.ExecContext(ctx, query, school.Name)
	if err != nil {
		return 0, err
	}

	uid, _ := result.LastInsertId()
	return uid, nil
}
