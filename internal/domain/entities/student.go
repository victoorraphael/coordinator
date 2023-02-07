package entities

import (
	"github.com/victoorraphael/coordinator/pkg/utils"
	"time"
)

// NewStudent returns a new Student as pointer
func NewStudent() *Student {
	return &Student{
		Person: Person{},
	}
}

type Student struct {
	ClassroomID int `db:"classroom_id"`
	SchoolID    int `db:"school_id"`
	Person
}

type StudentView struct {
	ClassroomUUID string `json:"classroom_uuid"`
	SchoolUUID    string `json:"school_uuid"`
	PersonView
}

type CreateStudent struct {
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Phone     string    `json:"phone"`
	Birthdate time.Time `json:"birthdate"`
	AddressID string    `json:"address_id"`
}

func (c CreateStudent) Validate() error {
	return utils.Validate(c)
}
