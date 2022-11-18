package student

import (
	"github.com/victoorraphael/school-plus-BE/internal/person"
)

type Student struct {
	person.Person
}

func (s Student) Add(t Student) (Student, error) {
	//TODO implement me
	panic("implement me")
}

func (s Student) Get(t Student) (Student, error) {
	return person.
}

func (s Student) List() ([]Student, error) {
	//TODO implement me
	panic("implement me")
}

func (s Student) Update(t Student) error {
	//TODO implement me
	panic("implement me")
}

func (s Student) Delete(t Student) error {
	//TODO implement me
	panic("implement me")
}

func New() Student {
	return Student{
		person.Person{
			Type: person.Student,
		},
	}
}
