package services

import (
	"context"
	"github.com/victoorraphael/coordinator/internal/adapters/repository"
	"github.com/victoorraphael/coordinator/internal/domain"
	"log"
)

func CreateAddress(ctx context.Context, repo repository.Address, addr *domain.Address) error {
	err := repo.Add(ctx, addr)
	if err != nil {
		log.Println("error:", err.Error())
	}

	return err
}
