package address

type Address struct {
	ID     int    `json:"-" db:"id"`
	Street string `json:"street" db:"street"`
	City   string `json:"city" db:"city"`
	Zip    string `json:"zip" db:"zip"`
	Number int    `json:"number" db:"number"`
}
