package repository

import (
	"context"
	"github.com/victoorraphael/coordinator/internal/adapters"
	"github.com/victoorraphael/coordinator/internal/domain"
	"github.com/victoorraphael/coordinator/internal/domain/contracts"
)

type address struct {
	pool adapters.DBPool
}

func NewAddressRepo(pool adapters.DBPool) contracts.AddressRepo {
	return &address{pool}
}

func (a *address) List() ([]domain.Address, error) {
	conn, err := a.pool.Acquire()
	if err != nil {
		return nil, err
	}
	defer a.pool.Release(conn)

	var resp []domain.Address
	_, errSelect := conn.Select("*").
		From("address").
		Load(&resp)

	return resp, errSelect
}

func (a address) Find(id int64) (domain.Address, error) {
	conn, err := a.pool.Acquire()
	if err != nil {
		return domain.Address{}, err
	}
	defer a.pool.Release(conn)

	resp := domain.Address{}
	_, err = conn.Select("street, city, zip, number").
		From("address").
		Where("id = $1", id).
		Load(&resp)

	return resp, err
}

func (a address) Add(addr *domain.Address) error {
	conn, err := a.pool.Acquire()
	if err != nil {
		return err
	}
	defer a.pool.Release(conn)

	return conn.
		InsertInto("address").
		Pair("street", addr.Street).
		Pair("city", addr.City).
		Pair("zip", addr.Zip).
		Pair("number", addr.Number).
		Returning("id").
		LoadContext(context.Background(), &addr.ID)
}
