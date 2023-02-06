package services

import (
	"context"
	"errors"
	"github.com/golangsugar/chatty"
	"github.com/victoorraphael/coordinator/internal/domain/entities"
	"github.com/victoorraphael/coordinator/internal/domain/repository"
	"github.com/victoorraphael/coordinator/pkg/uid"
)

type IAddressService interface {
	FetchAll(ctx context.Context) ([]entities.Address, error)
	Create(ctx context.Context, addr *entities.Address) error
	Update(ctx context.Context, addr entities.Address) error
}

type address struct {
	repo *repository.Repo
}

func NewAddressService(repo *repository.Repo) IAddressService {
	return &address{repo}
}

func (a *address) Update(ctx context.Context, addr entities.Address) error {
	return nil
}

func (a *address) FetchAll(ctx context.Context) ([]entities.Address, error) {
	return a.repo.Address.List(ctx)
}

func (a *address) Create(ctx context.Context, addr *entities.Address) error {
	exist, err := a.repo.Address.Search(ctx, entities.Address{
		Street: addr.Street,
		City:   addr.City,
		Zip:    addr.Zip,
	})
	if err != nil {
		chatty.Info(err.Error())
		return err
	}

	if exist.ID != 0 {
		return errors.New("endereço já existe")
	}

	addr.UUID = uid.NewUUID().String()
	return a.repo.Address.Add(ctx, addr)
}
