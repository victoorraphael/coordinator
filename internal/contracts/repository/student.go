package repository

import (
	"github.com/google/uuid"
	"github.com/victoorraphael/school-plus-BE/internal/entities"
)

type studentRepo struct {
	Adapters *entities.Adapters
	persons  IPersonRepo
}

type IStudentRepo interface {
	List() ([]entities.Student, error)
	FindOne(entities.Student) (entities.Student, error)
	Add(entities.Student) (uuid.UUID, error)
	Update(entities.Student) error
	Delete(entities.Student) error
}

func NewStudentRepo(adapters *entities.Adapters) IStudentRepo {
	return &studentRepo{
		Adapters: adapters,
		persons:  NewPersonRepo(adapters),
	}
}

func (sr studentRepo) List() ([]entities.Student, error) {
	persons, err := sr.persons.List()
	if err != nil {
		return nil, err
	}

	var students []entities.Student

	for _, p := range persons {
		students = append(students, entities.Student{
			ID:     p.ID,
			Person: p,
		})
	}

	return students, nil
}

func (sr studentRepo) FindOne(entities.Student) (entities.Student, error) {
	panic("implement me")
}

func (sr studentRepo) Add(std entities.Student) (uuid.UUID, error) {
	id, err := sr.persons.Add(std.Person)
	if err != nil {
		return uuid.Nil, err
	}

	sql := `
	INSERT INTO students (id)
	VALUES ($1)
	`

	errStd := sr.Adapters.DB.
		GetDatabase().
		QueryRow(sql, id).Err()

	if errStd != nil {
		return uuid.Nil, err
	}

	return id, nil
}

func (sr studentRepo) Update(entities.Student) error {
	panic("implement me")
}

func (sr studentRepo) Delete(entities.Student) error {
	panic("implement me")
}
