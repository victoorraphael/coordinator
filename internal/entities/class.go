package entities

import (
	"github.com/google/uuid"
	"github.com/victoorraphael/school-plus-BE/internal/student"
)

type Class struct {
	ID       uuid.UUID         `json:"id"`
	Name     string            `json:"name"`
	Subjects []Subject         `json:"subjects"`
	Students []student.Student `json:"students"`
}
