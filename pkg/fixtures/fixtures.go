package fixtures

import (
	_ "github.com/lib/pq" //
	"github.com/victoorraphael/coordinator/internal/domain/repository"
	"github.com/victoorraphael/coordinator/internal/domain/services"
	"github.com/victoorraphael/coordinator/pkg/database"
	"log"
)

type Adapters struct {
	Pool database.DBPool
	Repo *repository.Repo
	Srv  *services.Services
}

func (adapters *Adapters) Close() error {
	adapters.Pool.Close()
	adapters.Srv = nil
	adapters.Repo = nil
	return nil
}

func Connect() *Adapters {
	pool, err := database.NewPostgres(1)
	if err != nil {
		log.Fatal(err)
	}
	conn, err := pool.Acquire()
	if err != nil {
		log.Fatal(err)
	}

	if err := conn.Ping(); err != nil {
		log.Fatal(err)
	}
	pool.Release(conn)

	repo := repository.New(pool)
	srv := services.New(repo)

	return &Adapters{
		Pool: pool,
		Repo: repo,
		Srv:  srv,
	}
}
