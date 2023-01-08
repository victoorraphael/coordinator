package repository

import (
	"context"
	"github.com/victoorraphael/coordinator/internal/adapters/postgres"
	"github.com/victoorraphael/coordinator/internal/adapters/postgres/models"
	"github.com/victoorraphael/coordinator/internal/domain"
)

type Student struct{}

// Find all students from school
func (s *Student) Find(ctx context.Context, schoolID int) ([]domain.Student, error) {
	db := postgres.NewPostgresAdapter().GetDatabase()
	query := `
		SELECT p.id, p.uuid, p.name, p.email, p.phone, p.birthdate, p.type
		FROM students as s
		INNER JOIN persons as p on p.id = s.id 
		WHERE school_id = $1`

	rows, err := db.QueryContext(ctx, query, schoolID)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	response := make([]domain.Student, 0)
	for rows.Next() {
		item := &domain.Student{}
		err := rows.Scan(
			item.ID,
			item.UUID,
			item.Name,
			item.Email,
			item.Phone,
			item.Birthdate,
			item.Type,
		)
		if err != nil {
			return nil, err
		}

		response = append(response, *item)
	}

	return response, nil
}

func (s *Student) FindUUID(ctx context.Context, uuid string) (domain.Student, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Student) Create(ctx context.Context, student models.Student) (domain.Student, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Student) Update(ctx context.Context, student domain.Student) error {
	//TODO implement me
	panic("implement me")
}

func (s *Student) Delete(ctx context.Context, uuid string) error {
	//TODO implement me
	panic("implement me")
}
