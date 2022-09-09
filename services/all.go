package services

import (
	"github.com/victoorraphael/school-plus-BE/infra/entities"
	"github.com/victoorraphael/school-plus-BE/services/student"
)

func All(adapters *entities.Adapters) []entities.Service {
	return []entities.Service{
		student.New(adapters),
	}
}
