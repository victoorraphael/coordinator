package store

import (
	"github.com/victoorraphael/coordinator/internal/entities"
)

type Store struct {
	Student studentStore
}

// New returns an instance of Store with instances of all individuals stores
func New(adapters *entities.Adapters) *Store {
	return &Store{
		Student: studentStore{adapters: adapters},
	}
}
