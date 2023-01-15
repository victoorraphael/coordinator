package domain

import (
	"time"

	"github.com/google/uuid"
)

type Person struct {
	ID        int64      `json:"id"`
	UUID      uuid.UUID  `json:"uuid"`
	Name      string     `json:"name"`
	Email     string     `json:"email"`
	Phone     string     `json:"phone"`
	Birthdate time.Time  `json:"birthdate"`
	Type      PersonType `json:"type"`
	Address   Address    `json:"address"`
	School    School     `json:"school"`
}

// PersonType represents type of persons on system
type PersonType int

const (
	PersonStudent PersonType = iota + 1
	PersonProfessor
)

// String returns representative string of person type
func (t PersonType) String() string {
	names := []string{"student", "professor"}
	return names[t-1]
}
