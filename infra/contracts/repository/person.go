package repository

import (
	"github.com/google/uuid"
	"github.com/victoorraphael/school-plus-BE/infra/entities"
)

type personRepo struct {
	Adapters *entities.Adapters
}

type IPersonRepo interface {
	List() ([]entities.Person, error)
	FindOne(entities.Person) (entities.Person, error)
	Add(entities.Person) (uuid.UUID, error)
	Update(entities.Person) error
	Delete(entities.Person) error
}

func NewPersonRepo(adapters *entities.Adapters) IPersonRepo {
	return &personRepo{Adapters: adapters}
}

func (p personRepo) List() ([]entities.Person, error) {
	sql := `
	SELECT name, email, phone
	FROM persons
	`

	rows, err := p.Adapters.DB.
		GetDatabase().
		Query(sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	response := make([]entities.Person, 0)

	for rows.Next() {
		var p entities.Person

		err := rows.Scan(&p.ID, &p.Name, &p.Email, &p.Phone)
		if err != nil {
			continue
		}

		response = append(response, p)
	}

	return response, nil
}

func (p personRepo) FindOne(person entities.Person) (entities.Person, error) {
	panic("implement me")
}

func (p personRepo) Add(person entities.Person) (uuid.UUID, error) {
	sql := `
	INSERT INTO persons (name, email, phone)
	VALUES ($1, $2, $3)
	RETURNING id
	`
	var response entities.Person

	err := p.Adapters.DB.
		GetDatabase().
		QueryRow(sql, person.Name, person.Email, person.Phone).
		Scan(&response.ID)

	if err != nil {
		return uuid.Nil, err
	}

	return response.ID, nil
}

func (p personRepo) Update(person entities.Person) error {
	panic("implement me")
}

func (p personRepo) Delete(person entities.Person) error {
	panic("implement me")
}
