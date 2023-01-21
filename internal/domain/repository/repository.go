package repository

import (
	"github.com/victoorraphael/coordinator/pkg/database"
)

type Repo struct {
	Address IAddressRepository
	Person  IPersonRepository
	User    IUserRepository
}

func New(pool database.DBPool) *Repo {
	return &Repo{
		Address: NewAddressRepo(pool),
		Person:  NewPersonRepo(pool),
		User:    NewUserRepo(pool),
	}
}
