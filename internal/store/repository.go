package store

import (
	"errors"
	"github.com/victoorraphael/school-plus-BE/internal/entities"
)

type Repository struct {
	Adapters *entities.Adapters
}

func NewRepository(adapters *entities.Adapters) *Repository {
	return &Repository{Adapters: adapters}
}

var (
	ErrTypeNotEmpty = errors.New("type of person cannot be empty")
)

func (p Repository) List(tp entities.TypePerson) ([]entities.Person, error) {
	if tp == entities.Unknown {
		return nil, ErrTypeNotEmpty
	}

	sql := `
	SELECT id, uuid, name, email, phone, birthdate, created_at, type
	FROM persons
	WHERE type = $1
	`

	rows, err := p.Adapters.DB.
		GetDatabase().
		Query(sql, tp)
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

func (p Repository) FindOne(person entities.Person) (entities.Person, error) {
	panic("implement me")
}

func (p Repository) Add(person entities.Person) (int64, error) {
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
		return 0, err
	}

	return response.ID, nil
}

func (p Repository) Update(person entities.Person) error {
	panic("implement me")
}

func (p Repository) Delete(person entities.Person) error {
	panic("implement me")
}
