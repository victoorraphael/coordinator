package subject

import (
	"github.com/google/uuid"
	"github.com/victoorraphael/coordinator/internal/professor"
)

type Subject struct {
	ID        uuid.UUID           `json:"id"`
	Name      string              `json:"name"`
	Professor professor.Professor `json:"professor"`
}
