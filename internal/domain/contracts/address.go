package contracts

import (
	"github.com/victoorraphael/coordinator/internal/domain"
)

type AddressRepo interface {
	List() ([]domain.Address, error)
	Find(id int64) (domain.Address, error)
	Add(addr *domain.Address) error
}

type AddressService interface {
	FetchAll() ([]domain.Address, error)
	Create(*domain.Address) error
}
