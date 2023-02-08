package services

import (
	"github.com/victoorraphael/coordinator/internal/domain/repository"
)

type Services struct {
	Address IAddressService
	Person  IPersonService
	Auth    IAuthenticationService
	School  ISchoolService
	Subject ISubjectService
}

// New returns instances of all services
func New(repo *repository.Repo) *Services {
	return &Services{
		Address: NewAddressService(repo),
		Person:  NewPersonService(repo),
		Auth:    NewAuthenticationService(repo),
		School:  NewSchoolService(repo),
		Subject: NewSubjectService(repo),
	}
}
