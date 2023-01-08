package address

import (
	"context"
	"fmt"
	"github.com/victoorraphael/coordinator/internal/adapters"
	"github.com/victoorraphael/coordinator/internal/domain"
)

type Store struct {
	Adapters *adapters.Adapters
}

func (s Store) Find(ctx context.Context, field string, data any) (domain.Address, error) {
	db := s.Adapters.DB.GetDatabase()
	query := fmt.Sprintf("SELECT street, city, zip, number FROM address WHERE %s = $1", field)
	resp := domain.Address{}
	err := db.QueryRowContext(ctx, query, data).
		Scan(&resp.Street, &resp.City, &resp.Zip, &resp.Number)

	return resp, err
}

func (s Store) Create(ctx context.Context, address *domain.Address) error {
	db := s.Adapters.DB.GetDatabase()
	query := "INSERT INTO address (street, city, zip, number) VALUES ($1, $2, $3, $4) RETURNING id"
	return db.QueryRowContext(ctx, query, address.Street, address.City, address.Zip, address.Number).Scan(&address.ID)
}
