package models

type Address struct {
	ID     int    `db:"id"`
	Street string `db:"street"`
	City   string `db:"city"`
	ZIP    string `db:"zip"`
	Number int    `db:"number"`
}
