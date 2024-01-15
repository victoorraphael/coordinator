package entities

import (
	"time"
)

type Person struct {
	ID        int64      `db:"id" json:"id"`
	UUID      string     `db:"uuid" json:"uuid"`
	Name      string     `db:"name" json:"name"`
	Email     string     `db:"email" json:"email"`
	Phone     string     `db:"phone" json:"phone"`
	Birthdate time.Time  `db:"birthdate" json:"birthdate"`
	CreatedAt time.Time  `db:"created_at" json:"created_at"`
	Type      PersonType `db:"type" json:"type"`
	AddressID int64      `db:"address_id" json:"address_id"`
}

type PersonView struct {
	UUID        string    `json:"uuid"`
	Name        string    `json:"name"`
	Email       string    `json:"email"`
	Phone       string    `json:"phone"`
	Birthdate   time.Time `json:"birthdate"`
	AddressUUID string    `json:"address"`
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
