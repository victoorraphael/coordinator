package domain

import (
	"github.com/google/uuid"
)

type Classroom struct {
	ID   uuid.UUID
	Name string
	//Subjects []subject.Subject `json:"subjects"`
	Students []Student
}
