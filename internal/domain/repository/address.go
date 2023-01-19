package repository

import (
	"context"
	"github.com/victoorraphael/coordinator/internal/domain/entities"
	"github.com/victoorraphael/coordinator/pkg/database"
)

type IAddressRepository interface {
	List() ([]entities.Address, error)
	Find(id int64) (entities.Address, error)
	Add(addr *entities.Address) error
}

type address struct {
	pool database.DBPool
}

// NewAddressRepo returns a new IAddressRepository
func NewAddressRepo(pool database.DBPool) IAddressRepository {
	return &address{pool}
}

func (a *address) List() ([]entities.Address, error) {
	conn, err := a.pool.Acquire()
	if err != nil {
		return nil, err
	}
	defer a.pool.Release(conn)

	var resp []entities.Address
	_, errSelect := conn.Select("*").
		From("address").
		Load(&resp)

	return resp, errSelect
}

func (a address) Find(id int64) (entities.Address, error) {
	conn, err := a.pool.Acquire()
	if err != nil {
		return entities.Address{}, err
	}
	defer a.pool.Release(conn)

	resp := entities.Address{}
	_, err = conn.Select("street, city, zip, number").
		From("address").
		Where("id = ?", id).
		Load(&resp)

	return resp, err
}

func (a address) Add(addr *entities.Address) error {
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
