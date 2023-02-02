package entities

type School struct {
	ID        int64  `json:"-" db:"id"`
	UUID      string `json:"uuid" db:"uuid"`
	Name      string `db:"name"`
	AddressID int64  `db:"address_id"`
}
