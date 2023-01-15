package services

import (
	"github.com/victoorraphael/coordinator/internal/adapters/repository"
	"github.com/victoorraphael/coordinator/internal/domain"
	"github.com/victoorraphael/coordinator/internal/domain/contracts"
	"log"
)

type address struct {
	repo *repository.Repo
}

func NewAddressService(repo *repository.Repo) contracts.AddressService {
	return &address{repo}
}

func (a *address) FetchAll() ([]domain.Address, error) {
	return a.repo.Address.List()
}

func (a *address) Create(addr *domain.Address) error {
	err := a.repo.Address.Add(addr)
	if err != nil {
		log.Println("error:", err.Error())
	}

	return err
}
