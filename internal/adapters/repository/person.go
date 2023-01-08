package repository

import (
	"context"
	"github.com/google/uuid"
	"github.com/victoorraphael/coordinator/internal/adapters/postgres"
	"github.com/victoorraphael/coordinator/internal/adapters/postgres/models"
	"github.com/victoorraphael/coordinator/internal/domain"
	"log"
)

type Person struct{}

// List all persons from type domain.Person.Type
func (p *Person) List(ctx context.Context, person models.Person) ([]domain.Person, error) {
	db := postgres.NewPostgresAdapter().GetDatabase()
	query := "SELECT uuid, name, email, phone, birthdate, type FROM persons WHERE type = $1"
	rows, err := db.QueryContext(ctx, query, person.Type)
	if err != nil {
		log.Println("ERROR", err)
		return nil, err
	}
	defer rows.Close()

	persons := make([]domain.Person, 0)
	for rows.Next() {
		person := domain.Person{}
		if err := rows.Scan(&person.UUID, &person.Name, &person.Email, &person.Phone, &person.Birthdate,
			&person.Type); err != nil {
			return nil, err
		}

		persons = append(persons, person)
	}

	return persons, nil
}

// Add a new person
func (p *Person) Add(ctx context.Context, person models.Person) (uuid.UUID, error) {
	db := postgres.NewPostgresAdapter().GetDatabase()
	var respUUID uuid.UUID
	query := "INSERT INTO persons (name, email, phone, birthdate, type, address_id) VALUES ($1, $2, $3, $4, $5, $6) RETURNING uuid"
	err := db.QueryRowContext(ctx, query, person.Name, person.Email, person.Phone, person.Birthdate, person.Type,
		person.AddressID).
		Scan(&respUUID)

	return respUUID, err
}

// Delete a person based on uuid
func (p *Person) Delete(ctx context.Context, person models.Person) error {
	db := postgres.NewPostgresAdapter().GetDatabase()
	_, err := db.ExecContext(ctx, "DELETE FROM persons WHERE uuid = $1", person.UUID)
	return err
}

// Update a person based on uuid
func (p *Person) Update(ctx context.Context, person models.Person) error {
	db := postgres.NewPostgresAdapter().GetDatabase()
	query := "UPDATE persons SET name = $1, email = $2, phone = $3 WHERE uuid = $4"
	_, err := db.ExecContext(ctx, query, person.Name, person.Email, person.Phone, person.UUID)
	return err
}
