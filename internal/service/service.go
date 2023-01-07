package service

import (
	"github.com/victoorraphael/coordinator/internal/adapters"
	"github.com/victoorraphael/coordinator/internal/student"
)

type Service interface {
	Student() student.Service
}

type service struct {
	student student.Service
}

func New(adapters *adapters.Adapters) Service {
	return &service{student: student.NewService(adapters)}
}

func (s service) Student() student.Service { return s.student }
