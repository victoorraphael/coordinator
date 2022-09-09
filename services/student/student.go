package student

import (
	"errors"

	"github.com/victoorraphael/school-plus-BE/infra/contracts"
	"github.com/victoorraphael/school-plus-BE/infra/contracts/repository"
	"github.com/victoorraphael/school-plus-BE/infra/entities"
)

var (
	ErrMissingID = errors.New("missing ID")
)

type service struct {
	studentsRead  repository.IReadRepository[entities.Student]
	studentsWrite repository.IWriteRepository[entities.Student]
}

func New(adapters *entities.Adapters) contracts.IService[service] {

}
