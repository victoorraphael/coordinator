package service

import "github.com/victoorraphael/school-plus-BE/internal/entities"

type Service interface {
	StudentSRV() IStudentSRV
}

type service struct {
	student IStudentSRV
}

func New(adapters *entities.Adapters) Service {
	return &service{student: NewStudentService(adapters)}
}

func (s service) StudentSRV() IStudentSRV { return s.StudentSRV() }
