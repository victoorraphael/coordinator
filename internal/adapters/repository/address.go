package repository

import (
	"context"
	"fmt"
	"github.com/victoorraphael/coordinator/internal/adapters/postgres"
	"github.com/victoorraphael/coordinator/internal/domain"
)

type Address struct{}

func (a Address) Find(ctx context.Context, id int) (domain.Address, error) {
	db := postgres.NewPostgresAdapter().GetDatabase()
	query := fmt.Sprintf("SELECT street, city, zip, number FROM address WHERE id = $1")
	resp := domain.Address{}
	err := db.QueryRowContext(ctx, query, id).
		Scan(&resp.Street, &resp.City, &resp.Zip, &resp.Number)

	return resp, err
}

func (a Address) Create(ctx context.Context, address *domain.Address) error {
	db := postgres.NewPostgresAdapter().GetDatabase()
	query := "INSERT INTO address (street, city, zip, number) VALUES ($1, $2, $3, $4) RETURNING id"
	return db.QueryRowContext(ctx, query, address.Street, address.City, address.Zip, address.Number).Scan(&address.ID)
}
