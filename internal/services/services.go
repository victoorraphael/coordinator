package services

import (
	"github.com/victoorraphael/coordinator/internal/adapters/repository"
	"github.com/victoorraphael/coordinator/internal/domain/contracts"
)

type Services struct {
	Address contracts.AddressService
}

// New returns instances of all services
func New(repo *repository.Repo) *Services {
	return &Services{
		Address: NewAddressService(repo),
	}
}
