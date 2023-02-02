package fixtures

import (
	_ "github.com/lib/pq" //
	"github.com/victoorraphael/coordinator/internal/domain/repository"
	"github.com/victoorraphael/coordinator/internal/domain/services"
	"github.com/victoorraphael/coordinator/pkg/database"
	"log"
	"os"
)

type Adapters struct {
	Pool database.DBPool
	Repo *repository.Repo
	Srv  *services.Services
}

func (a *Adapters) Close() error {
	a.Pool.Close()
	a.Srv = nil
	a.Repo = nil
	return nil
}

func Connect() *Adapters {
	err := os.Setenv("DB_URI", "postgres://root:secret@abobrinha:5432/schoolplus?sslmode=disable")
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
		Pool: pool,
		Repo: repo,
		Srv:  srv,
	}
}
