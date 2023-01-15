package repository

import (
	"github.com/victoorraphael/coordinator/internal/adapters"
	"github.com/victoorraphael/coordinator/internal/domain"
)

type Address struct {
	pool adapters.DBPool
}

func (a *Address) List() ([]domain.Address, error) {
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

func (a Address) Find(id int64) (domain.Address, error) {
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

func (a Address) Add(addr *domain.Address) error {
	conn, err := a.pool.Acquire()
	if err != nil {
		return err
	}
	defer a.pool.Release(conn)

	_, errInsert := conn.InsertInto("address").
		Columns("street", "city", "zip", "number").
		Record(addr).
		Exec()

	return errInsert
}
