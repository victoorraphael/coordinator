package professor

import (
	"github.com/victoorraphael/coordinator/internal/person"
	"github.com/victoorraphael/coordinator/internal/school"
)

type Professor struct {
	School school.School `json:"school"`
	person.Person
}
