package services

import (
	"github.com/victoorraphael/coordinator/internal/domain/repository"
)

type Services struct {
	Address IAddressService
}

// New returns instances of all services
func New(repo *repository.Repo) *Services {
	return &Services{
		Address: NewAddressService(repo),
	}
}
