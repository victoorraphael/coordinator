package entities

import (
	"github.com/google/uuid"
	"github.com/victoorraphael/school-plus-BE/internal/person"
)

type Professor struct {
	ID uuid.UUID `json:"id"`
	person.Person
	Specialization string `json:"specialization"`
}
