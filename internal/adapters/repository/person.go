package repository

import (
	"github.com/victoorraphael/coordinator/internal/adapters"
	"github.com/victoorraphael/coordinator/internal/domain"
)

type Person struct{}

func (p *Person) List() []domain.Person {
	db := adapters.NewPostgresAdapter().GetDatabase()
}
