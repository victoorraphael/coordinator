package domain

import (
	"github.com/google/uuid"
)

type School struct {
	ID      uuid.UUID
	Address Address
	Classes []Classroom
}
