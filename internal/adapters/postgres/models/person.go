package models

import (
	"github.com/google/uuid"
	"github.com/victoorraphael/coordinator/internal/domain"
	"time"
)

type Person struct {
	ID        int64             `db:"id"`
	UUID      uuid.UUID         `db:"uuid"`
	Name      string            `db:"name"`
	Email     string            `db:"email"`
	Phone     string            `db:"phone"`
	Type      domain.PersonType `db:"type"`
	AddressID int64             `db:"address_id"`
	Birthdate time.Time         `db:"birthdate"`
	CreatedAt time.Time         `db:"created_at"`
}

func (p *Person) FromDomain(person domain.Person) {
	p.ID = person.ID
	p.UUID = person.UUID
	p.Name = person.Name
	p.Email = person.Email
	p.Phone = person.Phone
	p.Type = person.Type
	p.AddressID = person.Address.ID
	p.Birthdate = person.Birthdate
}
