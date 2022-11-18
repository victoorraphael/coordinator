package adapters

import (
	"database/sql"
	"errors"
	"log"
	"os"

	_ "github.com/lib/pq"
)

type PostgresAdapater struct {
	DB *sql.DB
}

func (p *PostgresAdapater) Connect() error {
	connStr := os.Getenv("DB_URI")

	if connStr == "" {
		return errors.New("empty mongo db connection string")
	}

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Println(err)
		return err
	}

	p.DB = db
	return nil
}

func (p *PostgresAdapater) Ping() bool {
	log.Println("trying to ping database...")
	err := p.DB.Ping()
	if err != nil {
		log.Println("failed to ping db, err:", err)
		return false
	}

	return true
}

func (p *PostgresAdapater) GetDatabase() *sql.DB {
	return p.DB
}
