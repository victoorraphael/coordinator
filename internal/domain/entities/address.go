package entities

type Address struct {
	ID     int64  `json:"-" db:"id"`
	UUID   string `json:"uuid" db:"uuid"`
	Street string `json:"street" db:"street"`
	City   string `json:"city" db:"city"`
	Zip    string `json:"zip" db:"zip"`
	Number int64  `json:"number" db:"number"`
}
