package store

import (
	"context"
	"fmt"
	"github.com/victoorraphael/coordinator/internal/entities"
	"log"
)

type personStore struct {
	adapters *entities.Adapters
}

func (s *personStore) List(ctx context.Context, person entities.Person) ([]entities.Person, error) {
	db := s.adapters.DB.GetDatabase()
	query := "SELECT uuid, name, email, phone, birthdate, type FROM persons WHERE type = $1"
	rows, err := db.QueryContext(ctx, query, person.Type)
	if err != nil {
		log.Println("ERROR", err)
		return nil, err
	}
	defer rows.Close()

	persons := make([]entities.Person, 0)
	for rows.Next() {
		person := entities.Person{}
		if err := rows.Scan(&person.UUID, &person.Name, &person.Email, &person.Phone, &person.Birthdate,
			&person.Type); err != nil {
			return nil, err
		}

		persons = append(persons, person)
	}

	return persons, nil
}

func (s *personStore) Add(ctx context.Context, person entities.Person) (entities.Person, error) {
	db := s.adapters.DB.GetDatabase()
	query := "INSERT INTO persons (name, email, phone, birthdate, type) VALUES ($1, $2, $3, $4, $5) RETURNING uuid"
	err := db.QueryRowContext(ctx, query, person.Name, person.Email, person.Phone, person.Birthdate, person.Type).
		Scan(&person.UUID)
	if err != nil {
		return entities.Person{}, err
	}

	return person, nil
}

func (s *personStore) FindByField(ctx context.Context, field string, value any) (entities.Person, error) {
	db := s.adapters.DB.GetDatabase()
	query := fmt.Sprintf("SELECT uuid, name, email, phone, birthdate, type FROM persons WHERE %s = $1", field)
	res := entities.Person{}
	err := db.QueryRowContext(ctx, query, value).
		Scan(&res.UUID, &res.Name, &res.Email, &res.Phone, &res.Birthdate, &res.Type)

	return res, err
}

func (s *personStore) Delete(ctx context.Context, student entities.Student) error {
	db := s.adapters.DB.GetDatabase()
	_, err := db.ExecContext(ctx, "DELETE FROM persons WHERE uuid = $1", student.UUID)
	return err
}
