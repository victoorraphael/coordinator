package services

import (
	"context"
	"github.com/victoorraphael/coordinator/internal/domain/entities"
	"github.com/victoorraphael/coordinator/internal/domain/repository"
	"github.com/victoorraphael/coordinator/pkg/uid"
)

type IAddressService interface {
	FetchAll(ctx context.Context) ([]entities.Address, error)
	Create(ctx context.Context, addr *entities.Address) error
}

type address struct {
	repo *repository.Repo
}

func NewAddressService(repo *repository.Repo) IAddressService {
	return &address{repo}
}

func (a *address) FetchAll(ctx context.Context) ([]entities.Address, error) {
	return a.repo.Address.List(ctx)
}

func (a *address) Create(ctx context.Context, addr *entities.Address) error {
	addr.UUID = uid.NewUUID().String()
	return a.repo.Address.Add(ctx, addr)
}
