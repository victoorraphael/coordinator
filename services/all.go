package services

import (
	"github.com/victoorraphael/school-plus-BE/infra/entities"
	"github.com/victoorraphael/school-plus-BE/services/student"
)

func All(adapters *entities.Adapters) []entities.Service {
	stdSrv := student.New(adapters)

	return []entities.Service{
		stdSrv,
	}
}
