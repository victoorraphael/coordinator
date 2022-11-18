package person

import (
	"github.com/google/uuid"
	"time"
)

type Person struct {
	ID        int64      `json:"id" db:"id"`
	UUID      uuid.UUID  `json:"uuid" db:"uuid"`
	Name      string     `json:"name" db:"name"`
	Email     string     `json:"email" db:"email"`
	Phone     string     `json:"phone" db:"phone"`
	Birthdate time.Time  `json:"birthdate" db:"birthdate"`
	Type      TypePerson `json:"type" db:"type"`
}

type TypePerson int

const (
	Unknown TypePerson = iota
	Student
	Professor
)

func (p TypePerson) String() string {
	names := []string{"student", "professor"}
	return names[p-1]
}
