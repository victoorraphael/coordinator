package fixtures

import (
	"context"
	"github.com/victoorraphael/coordinator/internal/domain/entities"
	"log"
)

var (
	addr      entities.Address
	classroom entities.Classroom
	school    entities.School
)

func Seed(adapters *Adapters) {
	ctx := context.Background()
	addr.Street = "rua teste"
	addr.City = "garanhuns"
	addr.Zip = "55295100"
	addr.Number = 123
	err := adapters.Repo.Address.Add(ctx, &addr)
	errFatal(err)

	classroom.Name = "school teste"
	errc := adapters.Repo.Classroom.Add(ctx, &classroom)
	errFatal(errc)

	school.Name = "school test"
	school.AddressID = addr.ID
	errs := adapters.Repo.School.Add(ctx, &school)
	errFatal(errs)
}

func errFatal(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
