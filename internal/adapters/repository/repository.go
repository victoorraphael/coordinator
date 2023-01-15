package repository

import (
	"github.com/victoorraphael/coordinator/internal/adapters"
	"github.com/victoorraphael/coordinator/internal/domain/contracts"
)

type Repo struct {
	Address contracts.AddressRepo
	//Person
	//School
	//Student
}

func New(pool adapters.DBPool) *Repo {
	return &Repo{
		Address: NewAddressRepo(pool),
		//Person:  Person{pool},
		//School:  School{pool},
		//Student: Student{pool},
	}
}
