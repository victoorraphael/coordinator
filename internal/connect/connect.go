package connect

import (
	"github.com/victoorraphael/school-plus-BE/internal/adapters"
	"github.com/victoorraphael/school-plus-BE/internal/entities"
)

func Connect() (*entities.Adapters, error) {
	postgres := &adapters.PostgresAdapater{}
	if err := postgres.Connect(); err != nil {
		return nil, err
	}

	return &entities.Adapters{
		DB: postgres,
	}, nil
}
