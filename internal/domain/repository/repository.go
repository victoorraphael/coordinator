package repository

import (
	"github.com/victoorraphael/coordinator/pkg/database"
)

type Repo struct {
	Address IAddressRepository
	//Person
	//School
	//Student
}

func New(pool database.DBPool) *Repo {
	return &Repo{
		Address: NewAddressRepo(pool),
		//Person:  Person{pool},
		//School:  School{pool},
		//Student: Student{pool},
	}
}
