package connect

import (
	"github.com/victoorraphael/coordinator/internal/adapters"
)

func Connect() (*adapters.Adapters, error) {
	return &adapters.Adapters{
		DB: adapters.NewPostgresAdapter(),
	}, nil
}
