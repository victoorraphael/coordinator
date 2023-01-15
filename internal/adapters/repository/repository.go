package repository

import "github.com/victoorraphael/coordinator/internal/adapters"

type Repo struct {
	Address
	Person
	School
	Student
}

func New(pool adapters.DBPool) *Repo {
	return &Repo{
		Address: Address{pool},
		//Person:  Person{pool},
		//School:  School{pool},
		//Student: Student{pool},
	}
}
