package entities

import (
	"time"

	"github.com/google/uuid"
)

type Person struct {
	ID        int64      `db:"id"`
	UUID      string     `db:"uuid"`
	Name      string     `db:"name"`
	Email     string     `db:"email"`
	Phone     string     `db:"phone"`
	Birthdate time.Time  `db:"birthdate"`
	CreatedAt time.Time  `db:"created_at"`
	Type      PersonType `db:"type"`
	AddressID int64      `db:"address_id"`
}

type PersonView struct {
	UUID      uuid.UUID `json:"uuid"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Phone     string    `json:"phone"`
	Birthdate time.Time `json:"birthdate"`
}

type PersonViewDetailed struct {
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
	PersonStudent PersonType = iota
	PersonProfessor
)

// String returns representative string of person type
func (t PersonType) String() string {
	names := []string{"student", "professor"}
	return names[t]
}
