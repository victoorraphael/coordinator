package services

import (
	"github.com/victoorraphael/coordinator/internal/adapters/repository"
	"github.com/victoorraphael/coordinator/internal/domain"
	"log"
)

type Address struct {
	repo *repository.Repo
}

func (a *Address) List() ([]domain.Address, error) {
	return a.repo.Address.List()
}

func (a *Address) CreateAddress(addr *domain.Address) error {
	err := a.repo.Address.Add(addr)
	if err != nil {
		log.Println("error:", err.Error())
	}

	return err
}
