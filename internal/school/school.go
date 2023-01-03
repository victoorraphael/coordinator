package school

import (
	"github.com/google/uuid"
	"github.com/victoorraphael/coordinator/internal/address"
	"github.com/victoorraphael/coordinator/internal/classroom"
)

type School struct {
	ID      uuid.UUID             `json:"id"`
	Address address.Address       `json:"address"`
	Classes []classroom.Classroom `json:"classrooms"`
}
