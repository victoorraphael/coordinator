package repository

import (
	"context"
	"github.com/victoorraphael/coordinator/internal/domain/entities"
	"github.com/victoorraphael/coordinator/pkg/database"
)

type IAddressRepository interface {
	List(ctx context.Context) ([]entities.Address, error)
	Find(ctx context.Context, id int64) (entities.Address, error)
	Add(ctx context.Context, addr *entities.Address) error
}

type address struct {
	pool database.DBPool
}

// NewAddressRepo returns a new IAddressRepository
func NewAddressRepo(pool database.DBPool) IAddressRepository {
	return &address{pool}
}

func (a *address) List(ctx context.Context) ([]entities.Address, error) {
	conn, err := a.pool.Acquire()
	if err != nil {
		return nil, err
	}
	defer a.pool.Release(conn)

	var resp []entities.Address
	_, errSelect := conn.Select("*").
		From("address").
		LoadContext(ctx, &resp)

	return resp, errSelect
}

func (a address) Find(ctx context.Context, id int64) (entities.Address, error) {
	conn, err := a.pool.Acquire()
	if err != nil {
		return entities.Address{}, err
	}
	defer a.pool.Release(conn)

	resp := entities.Address{}
	_, err = conn.Select("*").
		From("address").
		Where("id = ?", id).
		LoadContext(ctx, &resp)

	return resp, err
}

func (a address) Add(ctx context.Context, addr *entities.Address) error {
	conn, err := a.pool.Acquire()
	if err != nil {
		return err
	}
	defer a.pool.Release(conn)

	return conn.
		InsertInto("address").
		Pair("uuid", addr.UUID).
		Pair("street", addr.Street).
		Pair("city", addr.City).
		Pair("zip", addr.Zip).
		Pair("number", addr.Number).
		Returning("id").
		LoadContext(ctx, &addr.ID)
}
