package person

import (
	"context"
	"fmt"
	"github.com/victoorraphael/coordinator/internal/adapters"
	"github.com/victoorraphael/coordinator/internal/student"
	"log"
)

type Store struct {
	Adapters *adapters.Adapters
}

func (s Store) List(ctx context.Context, person Person) ([]Person, error) {
	db := s.Adapters.DB.GetDatabase()
	query := "SELECT uuid, name, email, phone, birthdate, type FROM persons WHERE type = $1"
	rows, err := db.QueryContext(ctx, query, person.Type)
	if err != nil {
		log.Println("ERROR", err)
		return nil, err
	}
	defer rows.Close()

	persons := make([]Person, 0)
	for rows.Next() {
		person := Person{}
		if err := rows.Scan(&person.UUID, &person.Name, &person.Email, &person.Phone, &person.Birthdate,
			&person.Type); err != nil {
			return nil, err
		}

		persons = append(persons, person)
	}

	return persons, nil
}

func (s Store) Add(ctx context.Context, person Person) (Person, error) {
	db := s.Adapters.DB.GetDatabase()
	query := "INSERT INTO persons (name, email, phone, birthdate, type, address_id) VALUES ($1, $2, $3, $4, $5, $6) RETURNING uuid"
	err := db.QueryRowContext(ctx, query, person.Name, person.Email, person.Phone, person.Birthdate, person.Type, person.Address.ID).
		Scan(&person.UUID)
	if err != nil {
		return Person{}, err
	}

	return person, nil
}

func (s Store) FindByField(ctx context.Context, field string, value any) (Person, error) {
	db := s.Adapters.DB.GetDatabase()
	query := fmt.Sprintf("SELECT uuid, name, email, phone, birthdate, type, address_id FROM persons WHERE %s = $1", field)
	res := Person{}
	err := db.QueryRowContext(ctx, query, value).
		Scan(&res.UUID, &res.Name, &res.Email, &res.Phone, &res.Birthdate, &res.Type, &res.Address.ID)

	return res, err
}

func (s Store) Delete(ctx context.Context, student student.Student) error {
	db := s.Adapters.DB.GetDatabase()
	_, err := db.ExecContext(ctx, "DELETE FROM persons WHERE uuid = $1", student.UUID)
	return err
}

func (s Store) Update(ctx context.Context, student student.Student) error {
	db := s.Adapters.DB.GetDatabase()
	query := "UPDATE persons SET name = $1, email = $2, phone = $3 WHERE uuid = $4"
	log.Println("PERSON", student)
	_, err := db.ExecContext(ctx, query, student.Name, student.Email, student.Phone, student.UUID)
	return err
}
