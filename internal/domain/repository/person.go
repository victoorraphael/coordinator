package repository

import (
	"context"
	"github.com/victoorraphael/coordinator/internal/domain/entities"
	"github.com/victoorraphael/coordinator/pkg/database"
)

type IPersonRepository interface {
	List(entities.PersonType) ([]entities.Person, error)
	Add(*entities.Person) error
	FindUUID(uuid string) (entities.Person, error)
	FindID(ctx context.Context, id int64) (entities.Person, error)
	Delete(uuid string) error
	Update(entities.Person) error
}

type person struct {
	pool database.DBPool
}

func (p *person) List(t entities.PersonType) ([]entities.Person, error) {
	conn, err := p.pool.Acquire()
	if err != nil {
		return nil, err
	}
	defer p.pool.Release(conn)

	var res []entities.Person
	_, errSelect := conn.Select("*").
		From("persons as p").
		Where("p.type = ?", int(t)).
		LoadContext(context.Background(), &res)

	return res, errSelect
}

func (p *person) Add(person *entities.Person) error {
	conn, err := p.pool.Acquire()
	if err != nil {
		return err
	}
	defer p.pool.Release(conn)

	return conn.
		InsertInto("persons").
		Pair("name", person.Name).
		Pair("email", person.Email).
		Pair("phone", person.Phone).
		Pair("birthdate", person.Birthdate).
		Pair("type", person.Type).
		Pair("address_id", person.AddressID).
		Returning("uuid").
		LoadContext(context.Background(), &person.UUID)
}

func (p *person) FindID(ctx context.Context, id int64) (entities.Person, error) {
	conn, err := p.pool.Acquire()
	if err != nil {
		return entities.Person{}, err
	}
	defer p.pool.Release(conn)

	var res entities.Person
	_, err = conn.Select("*").
		From("persons").
		Where("id = ?", id).
		LoadContext(ctx, &res)
	return res, err
}

func (p *person) FindUUID(uuid string) (entities.Person, error) {
	//TODO implement me
	panic("implement me")
}

func (p *person) Delete(uuid string) error {
	//TODO implement me
	panic("implement me")
}

func (p *person) Update(person entities.Person) error {
	//TODO implement me
	panic("implement me")
}

func NewPersonRepo(pool database.DBPool) IPersonRepository {
	return &person{pool}
}
