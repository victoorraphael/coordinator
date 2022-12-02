package service

import (
	"github.com/victoorraphael/coordinator/internal/store"
)

type Service interface {
	StudentSRV() IStudentSRV
}

type service struct {
	student IStudentSRV
}

func New(s *store.Store) Service {
	return &service{student: NewStudentService(s)}
}

func (s service) StudentSRV() IStudentSRV { return s.student }
