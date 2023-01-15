package domain

type Address struct {
	ID     int64  `json:"id,omitempty" db:"id"`
	Street string `json:"street,omitempty" db:"street"`
	City   string `json:"city,omitempty" db:"city"`
	Zip    string `json:"zip,omitempty" db:"zip"`
	Number int    `json:"number,omitempty" db:"number"`
}
