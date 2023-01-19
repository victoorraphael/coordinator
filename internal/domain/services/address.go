package services

import (
	"github.com/victoorraphael/coordinator/internal/domain/entities"
	"github.com/victoorraphael/coordinator/internal/domain/repository"
)

type IAddressService interface {
	FetchAll() ([]entities.Address, error)
	Create(*entities.Address) error
}

type address struct {
	repo *repository.Repo
}

func NewAddressService(repo *repository.Repo) IAddressService {
	return &address{repo}
}

func (a *address) FetchAll() ([]entities.Address, error) {
	return a.repo.Address.List()
}

func (a *address) Create(addr *entities.Address) error {
	return a.repo.Address.Add(addr)
}
