package entities

import (
	"github.com/google/uuid"
)

type Class struct {
	ID       uuid.UUID `json:"id"`
	Name     string    `json:"name"`
	Subjects []Subject `json:"subjects"`
	Students []Student `json:"students"`
}
