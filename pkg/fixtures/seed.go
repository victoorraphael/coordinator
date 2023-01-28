package fixtures

import (
	"github.com/victoorraphael/coordinator/internal/domain/entities"
	"log"
)

func Seed(adapters *Adapters) {
	err := adapters.repo.Address.Add(&entities.Address{
		Street: "rua teste",
		City:   "garanhuns",
		Zip:    "55295100",
		Number: 123,
	})
	errFatal(err)

}

func errFatal(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
