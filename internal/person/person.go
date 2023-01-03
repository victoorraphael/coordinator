package person

import (
	"github.com/google/uuid"
	"github.com/victoorraphael/coordinator/internal/address"
	"time"
)

type Person struct {
	ID        int64           `json:"id" db:"id"`
	UUID      uuid.UUID       `json:"uuid" db:"uuid"`
	Name      string          `json:"name" db:"name"`
	Email     string          `json:"email" db:"email"`
	Phone     string          `json:"phone" db:"phone"`
	Birthdate time.Time       `json:"birthdate" db:"birthdate"`
	CreatedAt time.Time       `json:"created_at" db:"created_at"`
	Type      Type            `json:"type" db:"type"`
	Address   address.Address `json:"address"`
}

// Type represents type of persons on system
type Type int

const (
	Student Type = iota + 1
	Professor
)

// String returns representative string of person type
func (p Type) String() string {
	names := []string{"student", "professor"}
	return names[p-1]
}
