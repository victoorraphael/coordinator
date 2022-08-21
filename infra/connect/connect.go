package connect

import (
	"github.com/victoorraphael/school-plus-BE/infra/adapters"
	"github.com/victoorraphael/school-plus-BE/infra/entities"
)

func Connect() (*entities.Adapters, error) {
	mongoAdapt := adapters.MongoAdapter{}
	err := mongoAdapt.Connect()
	if err != nil {
		return nil, err
	}

	return &entities.Adapters{
		DB: &mongoAdapt,
	}, nil
}
