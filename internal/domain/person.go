package domain

import (
	"time"

	"github.com/google/uuid"
)

type Person struct {
	ID        int64
	UUID      uuid.UUID
	Name      string
	Email     string
	Phone     string
	Birthdate time.Time
	Type      Type
	Address   Address
}

// Type represents type of persons on system
type Type int

const (
	PersonStudent Type = iota + 1
	PersonProfessor
)

// String returns representative string of person type
func (p Type) String() string {
	names := []string{"student", "professor"}
	return names[p-1]
}
