package entities

import "github.com/victoorraphael/coordinator/pkg/helpers"

type School struct {
	ID        int64  `db:"id"`
	UUID      string `db:"uuid"`
	Name      string `db:"name"`
	AddressID int64  `db:"address_id"`
}

type CreateSchool struct {
	Name        string `json:"name"`
	AddressUUID string `json:"address_uuid"`
}

func (s CreateSchool) Validate() error {
	return helpers.Validate(s)
}
