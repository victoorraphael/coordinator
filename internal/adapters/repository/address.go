package repository

import (
	"context"
	"fmt"
	"github.com/victoorraphael/coordinator/internal/adapters/postgres"
	"github.com/victoorraphael/coordinator/internal/domain"
)

type Address struct{}

func (a Address) Find(ctx context.Context, id int64) (domain.Address, error) {
	db := postgres.NewPostgresAdapter().GetDatabase()
	query := fmt.Sprintf("SELECT street, city, zip, number FROM address WHERE id = $1")
	resp := domain.Address{}
	err := db.QueryRowContext(ctx, query, id).
		Scan(&resp.Street, &resp.City, &resp.Zip, &resp.Number)

	return resp, err
}

func (a Address) Add(ctx context.Context, addr *domain.Address) error {
	db := postgres.NewPostgresAdapter().GetDatabase()
	query := "INSERT INTO address (street, city, zip, number) VALUES ($1, $2, $3, $4) RETURNING id"
	return db.QueryRowContext(ctx, query, addr.Street, addr.City, addr.Zip, addr.Number).Scan(&addr.ID)
}
