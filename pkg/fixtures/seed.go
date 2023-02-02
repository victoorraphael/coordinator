package fixtures

import (
	"context"
	"github.com/victoorraphael/coordinator/internal/domain/entities"
	"github.com/victoorraphael/coordinator/pkg/security"
	"github.com/victoorraphael/coordinator/pkg/uid"
	"log"
	"time"
)

var (
	addr      entities.Address
	classroom entities.Classroom
	school    entities.School
	person    entities.Person
)

func (adapters *Adapters) Seed() *Adapters {
	ctx := context.Background()
	addr.UUID = uid.NewUUID().String()
	addr.Street = "rua teste"
	addr.City = "garanhuns"
	addr.Zip = "55295100"
	addr.Number = 123
	err := adapters.Repo.Address.Add(ctx, &addr)
	errFatal(err)

	classroom.UUID = uid.NewUUID().String()
	classroom.Name = "school teste"
	errc := adapters.Repo.Classroom.Add(ctx, &classroom)
	errFatal(errc)

	school.UUID = uid.NewUUID().String()
	school.Name = "school test"
	school.AddressID = addr.ID
	errs := adapters.Repo.School.Add(ctx, &school)
	errFatal(errs)

	person.UUID = uid.NewUUID().String()
	person.Name = "demo"
	person.Email = "demo@email.com"
	person.Phone = "999998877"
	person.Birthdate = time.Now().AddDate(-30, 0, 0)
	person.Type = entities.PersonStudent
	person.AddressID = addr.ID
	errFatal(adapters.Repo.Person.Add(&person))
	pass, _ := security.HashPassword("supersecret")
	user := entities.User{
		PersonID:     person.ID,
		Email:        person.Email,
		PasswordHash: pass,
	}
	errFatal(adapters.Repo.User.Add(ctx, &user))

	return adapters
}

func errFatal(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
