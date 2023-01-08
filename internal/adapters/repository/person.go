package repository

import (
	"context"
	"fmt"
	"github.com/victoorraphael/coordinator/internal/adapters/postgres"
	"github.com/victoorraphael/coordinator/internal/domain"
	"log"
)

type Person struct{}

func (p *Person) List(ctx context.Context, person domain.Person) ([]domain.Person, error) {
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

func (p *Person) Add(ctx context.Context, person domain.Person) (domain.Person, error) {
	db := postgres.NewPostgresAdapter().GetDatabase()
	query := "INSERT INTO persons (name, email, phone, birthdate, type, address_id) VALUES ($1, $2, $3, $4, $5, $6) RETURNING uuid"
	err := db.QueryRowContext(ctx, query, person.Name, person.Email, person.Phone, person.Birthdate, person.Type, person.Address.ID).
		Scan(&person.UUID)
	if err != nil {
		return domain.Person{}, err
	}

	return person, nil
}

func (p *Person) FindByField(ctx context.Context, field string, value any) (domain.Person, error) {
	db := postgres.NewPostgresAdapter().GetDatabase()
	query := fmt.Sprintf("SELECT uuid, name, email, phone, birthdate, type, address_id FROM persons WHERE %s = $1", field)
	res := domain.Person{}
	err := db.QueryRowContext(ctx, query, value).
		Scan(&res.UUID, &res.Name, &res.Email, &res.Phone, &res.Birthdate, &res.Type, &res.Address.ID)

	return res, err
}

func (p *Person) Delete(ctx context.Context, student domain.Student) error {
	db := postgres.NewPostgresAdapter().GetDatabase()
	_, err := db.ExecContext(ctx, "DELETE FROM persons WHERE uuid = $1", student.UUID)
	return err
}

func (p *Person) Update(ctx context.Context, student domain.Student) error {
	db := postgres.NewPostgresAdapter().GetDatabase()
	query := "UPDATE persons SET name = $1, email = $2, phone = $3 WHERE uuid = $4"
	log.Println("PERSON", student)
	_, err := db.ExecContext(ctx, query, student.Name, student.Email, student.Phone, student.UUID)
	return err
}
