package repository_test

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/victoorraphael/coordinator/internal/domain/entities"
	"github.com/victoorraphael/coordinator/internal/domain/repository"
	"github.com/victoorraphael/coordinator/pkg/fixtures"
	"math/rand"
	"testing"
)

func TestAddress_Add(t *testing.T) {
	adapters := fixtures.Connect()
	repo := repository.NewAddressRepo(adapters.Pool)
	addr := &entities.Address{
		Street: "rua teste",
		City:   "cidade teste",
		Zip:    "1878372",
		Number: int64(rand.Intn(10000)),
	}
	err := repo.Add(context.Background(), addr)
	assert.Nil(t, err)
	assert.Greater(t, addr.ID, int64(0))
}

func TestAddress_List(t *testing.T) {
	adapters := fixtures.Connect()
	repo := repository.NewAddressRepo(adapters.Pool)
	list, err := repo.List(context.Background())
	assert.Nil(t, err)
	assert.Greater(t, len(list), 0)
}
