package entities

type Subject struct {
	ID   int64  `json:"-" db:"id"`
	UUID string `json:"uuid" db:"uuid"`
	Name string `json:"name" db:"nome"` //TODO: ajustar coluna no banco para "name"
}
