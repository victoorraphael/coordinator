package student

import (
	"github.com/victoorraphael/coordinator/internal/address"
	"github.com/victoorraphael/coordinator/internal/classroom"
	"github.com/victoorraphael/coordinator/internal/person"
	"github.com/victoorraphael/coordinator/internal/school"
)

type Student struct {
	Classroom classroom.Classroom `json:"classroom"`
	School    school.School       `json:"school"`
	person.Person
}

// NewStudent returns a new Student as pointer
func NewStudent(address address.Address) *Student {
	p := person.NewPersonBuilder().
		As(person.Student).
		LivesAt(address).
		Build()

	return &Student{
		Classroom: classroom.Classroom{},
		School:    school.School{},
		Person:    p,
	}
}
