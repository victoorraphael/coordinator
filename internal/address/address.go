package address

import "github.com/google/uuid"

type Address struct {
	ID     int       `json:"-" db:"id"`
	UUID   uuid.UUID `json:"uuid" db:"uuid"`
	Street string    `json:"street" db:"street"`
	City   string    `json:"city" db:"city"`
	Zip    string    `json:"zip" db:"zip"`
	Number int       `json:"number" db:"number"`
}
