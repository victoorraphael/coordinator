package store

import (
	"context"
	"github.com/victoorraphael/coordinator/internal/entities"
	"log"
)

type studentStore struct {
	adapters *entities.Adapters
}

func (s *studentStore) List(_ context.Context) ([]entities.Student, error) {
	db := s.adapters.DB.GetDatabase()
	query := "SELECT uuid, name, email, phone, birthdate FROM persons WHERE type = $1"
	rows, err := db.Query(query, entities.PersonStudent)
	if err != nil {
		log.Println("ERROR", err)
		return nil, err
	}
	defer rows.Close()

	students := make([]entities.Student, 0)
	for rows.Next() {
		student := entities.NewStudent()
		if err := rows.Scan(&student.UUID, &student.Name, &student.Email, &student.Phone, &student.Birthdate); err != nil {
			return nil, err
		}

		students = append(students, student)
	}

	return students, nil
}
