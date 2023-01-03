package professor

import (
	"github.com/google/uuid"
)

type Professor struct {
	ID uuid.UUID `json:"id"`
	Person
	Specialization string `json:"specialization"`
}
