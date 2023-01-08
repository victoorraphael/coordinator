package models

import (
	"github.com/google/uuid"
	"github.com/victoorraphael/coordinator/internal/domain"
	"time"
)

type Person struct {
	ID        int               `db:"id"`
	UUID      uuid.UUID         `db:"uuid"`
	Name      string            `db:"name"`
	Email     string            `db:"email"`
	Phone     string            `db:"phone"`
	Type      domain.PersonType `db:"type"`
	AddressID int               `db:"address_id"`
	Birthdate time.Time         `db:"birthdate"`
	CreatedAt time.Time         `db:"created_at"`
}
