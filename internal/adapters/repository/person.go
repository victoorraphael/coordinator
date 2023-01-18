package repository

import (
	"github.com/google/uuid"
	"github.com/victoorraphael/coordinator/internal/adapters"
	"github.com/victoorraphael/coordinator/internal/adapters/postgres/models"
	"github.com/victoorraphael/coordinator/internal/domain"
	"github.com/victoorraphael/coordinator/internal/domain/contracts"
)

type person struct {
	pool adapters.DBPool
}

func NewPersonRepo(pool adapters.DBPool) contracts.PersonRepo {
	return &person{pool}
}

func (p *person) List(person models.Person) ([]domain.Person, error) {
	conn, err := p.pool.Acquire()
	if err != nil {
		return nil, err
	}
	defer p.pool.Release(conn)

	var res []models.Person
	_, err = conn.Select("*").
		From("persons").
		Load(&res)
	return res, err
}

func (p *person) Add(person models.Person) (uuid.UUID, error) {
	//TODO implement me
	panic("implement me")
}
