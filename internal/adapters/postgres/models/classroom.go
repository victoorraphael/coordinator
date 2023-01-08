package models

type Classroom struct {
	ID   int    `db:"id"`
	Name string `db:"name"`
}
