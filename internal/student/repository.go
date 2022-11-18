package student

import (
	"github.com/victoorraphael/school-plus-BE/internal/entities"
	"github.com/victoorraphael/school-plus-BE/internal/person"
)

type Repository struct {
	Adapters *entities.Adapters
	persons  *person.Repository
}
