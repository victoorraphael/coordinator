package connect

import (
	"github.com/victoorraphael/coordinator/internal/adapters"
	"github.com/victoorraphael/coordinator/internal/entities"
)

func Connect() (*entities.Adapters, error) {
	return &entities.Adapters{
		DB: adapters.NewPostgresAdapter(),
	}, nil
}
