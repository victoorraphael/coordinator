package repository

import (
	"context"
	"github.com/victoorraphael/coordinator/internal/domain/entities"
	"github.com/victoorraphael/coordinator/pkg/database"
)

type IPersonRepository interface {
	List(ctx context.Context, t entities.PersonType) ([]entities.Person, error)
	Add(ctx context.Context, p *entities.Person) error
	FindUUID(ctx context.Context, uuid string) (entities.Person, error)
	FindID(ctx context.Context, id int64) (entities.Person, error)
	Delete(ctx context.Context, uuid string) error
	Update(ctx context.Context, p entities.Person) error
}

type person struct {
	pool database.DBPool
}

func (p *person) List(ctx context.Context, t entities.PersonType) ([]entities.Person, error) {
	conn, err := p.pool.Acquire()
	if err != nil {
		return nil, err
	}
	defer p.pool.Release(conn)

	var res []entities.Person
	_, errSelect := conn.Select("*").
		From("persons as p").
		Where("p.type = ?", int(t)).
		LoadContext(ctx, &res)

	return res, errSelect
}

func (p *person) Add(ctx context.Context, person *entities.Person) error {
	conn, err := p.pool.Acquire()
	if err != nil {
		return err
	}
	defer p.pool.Release(conn)

	return conn.
		InsertInto("persons").
		Pair("uuid", person.UUID).
		Pair("name", person.Name).
		Pair("email", person.Email).
		Pair("phone", person.Phone).
		Pair("birthdate", person.Birthdate).
		Pair("type", person.Type).
		Pair("address_id", person.AddressID).
		Returning("id").
		LoadContext(ctx, &person.ID)
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

func (p *person) FindUUID(ctx context.Context, uuid string) (entities.Person, error) {
	conn, err := p.pool.Acquire()
	if err != nil {
		return entities.Person{}, err
	}
	defer p.pool.Release(conn)

	var res entities.Person
	_, err = conn.Select("*").
		From("persons").
		Where("uuid = ?", uuid).
		LoadContext(ctx, &res)
	return res, err
}

func (p *person) Delete(ctx context.Context, uuid string) error {
	conn, err := p.pool.Acquire()
	if err != nil {
		return err
	}
	defer p.pool.Release(conn)

	_, err = conn.DeleteFrom("persons").
		Where("uuid = ?", uuid).
		ExecContext(ctx)
	return err
}

func (p *person) Update(ctx context.Context, person entities.Person) error {
	conn, err := p.pool.Acquire()
	if err != nil {
		return err
	}
	defer p.pool.Release(conn)

	var res entities.Person
	err = conn.Update("persons").
		Set("name", person.Name).
		Set("email", person.Email).
		Set("phone", person.Phone).
		Set("birthdate", person.Birthdate).
		Set("address_id", person.AddressID).
		Where("uuid = ?", person.UUID).
		LoadContext(ctx, &res)
	return err
}

func NewPersonRepo(pool database.DBPool) IPersonRepository {
	return &person{pool}
}
