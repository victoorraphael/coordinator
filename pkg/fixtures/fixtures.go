package fixtures

import (
	"github.com/victoorraphael/coordinator/internal/domain/repository"
	"github.com/victoorraphael/coordinator/internal/domain/services"
	"github.com/victoorraphael/coordinator/pkg/database"
	"log"
	"os"
)

type Adapters struct {
	pool database.DBPool
	repo *repository.Repo
	srv  *services.Services
}

func Connect() *Adapters {
	err := os.Setenv("DB_URI", "postgres://root:secret@localhost:5432/schoolplus?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
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
		pool: pool,
		repo: repo,
		srv:  srv,
	}
}
