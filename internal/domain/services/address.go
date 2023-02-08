package services

import (
	"context"
	"errors"
	"github.com/gocraft/dbr/v2"
	"github.com/golangsugar/chatty"
	"github.com/victoorraphael/coordinator/internal/domain/entities"
	"github.com/victoorraphael/coordinator/internal/domain/repository"
	"github.com/victoorraphael/coordinator/pkg/errs"
	"github.com/victoorraphael/coordinator/pkg/uid"
)

type IAddressService interface {
	FetchAll(ctx context.Context) ([]entities.Address, error)
	Create(ctx context.Context, addr *entities.Address) error
	Find(ctx context.Context, addr entities.Address) (entities.Address, error)
}

type address struct {
	repo *repository.Repo
}

func NewAddressService(repo *repository.Repo) IAddressService {
	return &address{repo}
}

func (a *address) Find(ctx context.Context, addr entities.Address) (entities.Address, error) {
	if addr.UUID == "" && addr.ID <= 0 {
		return entities.Address{}, errs.WrapError(errs.ErrInternalError, "missing query fields uuid")
	}

	found, err := a.repo.Address.Search(ctx, addr)
	if err != nil {
		return entities.Address{}, err
	}

	return found, nil
}

func (a *address) FetchAll(ctx context.Context) ([]entities.Address, error) {
	list, err := a.repo.Address.List(ctx)
	if err != nil {
		return nil, errs.WrapError(errs.ErrInternalError, "nao foi possivel buscar enderecos")
	}
	return list, nil
}

func (a *address) Create(ctx context.Context, addr *entities.Address) error {
	exist, err := a.repo.Address.Search(ctx, entities.Address{
		Street: addr.Street,
		City:   addr.City,
		Zip:    addr.Zip,
	})
	if err != nil && !errors.Is(err, dbr.ErrNotFound) {
		chatty.Errorf("erro ao buscar endereco %#v: err: %v", *addr, err)
		return errs.WrapError(errs.ErrInternalError, "falha ao criar endereco")
	}

	if exist.ID != 0 {
		return errs.WrapError(errs.ErrFieldViolation, "endereco ja cadastrado")
	}

	addr.UUID = uid.NewUUID().String()
	if err := a.repo.Address.Add(ctx, addr); err != nil {
		chatty.Errorf("falha ao criar endereco: err: %v", err)
		return errs.WrapError(errs.ErrInternalError, "nao foi possivel criar endereco")
	}
	return nil
}
