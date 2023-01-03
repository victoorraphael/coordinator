package store

import (
	"github.com/victoorraphael/coordinator/internal/adapters"
)

type Store struct {
	Student studentStore
}

// New returns an instance of Store with instances of all individuals stores
func New(adapters *adapters.Adapters) *Store {
	person := personStore{adapters: adapters}
	return &Store{
		Student: studentStore{adapters: adapters, person: person},
	}
}
