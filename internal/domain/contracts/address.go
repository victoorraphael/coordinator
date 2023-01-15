package contracts

import (
	"github.com/victoorraphael/coordinator/internal/adapters"
	"github.com/victoorraphael/coordinator/internal/domain"
)

type IAddress interface {
	Find(pool adapters.DBPool, id int64) (domain.Address, error)
	Add(pool adapters.DBPool, addr *domain.Address) error
}
