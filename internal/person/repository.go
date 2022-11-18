package person

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

func (p Repository) List(tp TypePerson) ([]Person, error) {
	if tp == Unknown {
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

	response := make([]Person, 0)

	for rows.Next() {
		var p Person

		err := rows.Scan(&p.ID, &p.Name, &p.Email, &p.Phone)
		if err != nil {
			continue
		}

		response = append(response, p)
	}

	return response, nil
}

func (p Repository) FindOne(person Person) (Person, error) {
	panic("implement me")
}

func (p Repository) Add(person Person) (int64, error) {
	sql := `
	INSERT INTO persons (name, email, phone)
	VALUES ($1, $2, $3)
	RETURNING id
	`
	var response Person

	err := p.Adapters.DB.
		GetDatabase().
		QueryRow(sql, person.Name, person.Email, person.Phone).
		Scan(&response.ID)

	if err != nil {
		return 0, err
	}

	return response.ID, nil
}

func (p Repository) Update(person Person) error {
	panic("implement me")
}

func (p Repository) Delete(person Person) error {
	panic("implement me")
}
