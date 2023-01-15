package services

import "github.com/victoorraphael/coordinator/internal/adapters/repository"

type Services struct {
	Address
}

func New(repo *repository.Repo) *Services {
	return &Services{
		Address{repo},
	}
}
