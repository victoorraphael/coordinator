package entities

import (
	"github.com/google/uuid"
)

type Student struct {
	ID uuid.UUID `json:"id"`
	Person
}
