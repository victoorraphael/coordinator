package store

import (
	"context"
	"github.com/google/uuid"
	"github.com/victoorraphael/coordinator/internal/entities"
	"log"
	"time"
)

type personStore struct {
	adapters *entities.Adapters
}

func (s *personStore) List(_ context.Context, person entities.Person) ([]entities.Person, error) {
	db := s.adapters.DB.GetDatabase()
	query := "SELECT uuid, name, email, phone, birthdate, type FROM persons WHERE type = $1"
	rows, err := db.Query(query, person.Type)
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

func (s *personStore) Add(_ context.Context, person entities.Person) (entities.Person, error) {
	db := s.adapters.DB.GetDatabase()
	bday, _ := time.Parse(time.RFC3339, person.Birthdate)
	query := "INSERT INTO persons (name, email, phone, birthdate, type) VALUES ($1, $2, $3, $4, $5) RETURNING uuid"
	var uuidQuery string
	err := db.QueryRow(query, person.Name, person.Email, person.Phone, bday, person.Type).Scan(&uuidQuery)
	if err != nil {
		return entities.Person{}, err
	}

	uid := uuid.MustParse(uuidQuery)
	return entities.Person{UUID: uid}, nil
}
