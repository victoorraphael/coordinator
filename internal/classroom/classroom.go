package classroom

import (
	"github.com/google/uuid"
	"github.com/victoorraphael/coordinator/internal/student"
	"github.com/victoorraphael/coordinator/internal/subject"
)

type Classroom struct {
	ID       uuid.UUID         `json:"id"`
	Name     string            `json:"name"`
	Subjects []subject.Subject `json:"subjects"`
	Students []student.Student `json:"students"`
}
