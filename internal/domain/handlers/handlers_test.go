package handlers_test

import (
	"github.com/victoorraphael/coordinator/internal/domain/repository"
	"github.com/victoorraphael/coordinator/internal/domain/services"
	"github.com/victoorraphael/coordinator/pkg/database"
	"log"
	"os"
	"testing"
)

var (
	dbPool database.DBPool
	srvs   *services.Services
)

func TestMain(m *testing.M) {
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
	dbPool = pool

	repo := repository.New(dbPool)
	srvs = services.New(repo)

	code := m.Run()

	dbPool.Close()
	os.Exit(code)
}
