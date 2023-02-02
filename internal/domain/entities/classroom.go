package entities

type Classroom struct {
	ID   int64  `json:"id"`
	UUID string `json:"uuid" db:"uuid"`
	Name string `json:"name"`
}
