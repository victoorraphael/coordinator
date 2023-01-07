package adapters

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
)

type PostgresAdapter struct {
	db *sql.DB
}

func NewPostgresAdapter() DBAdapter {
	p := &PostgresAdapter{}
	p.connect()
	return p
}

// connect try to connect DB with environment variable
func (p *PostgresAdapter) connect() {
	connStr := os.Getenv("DB_URI")

	if connStr == "" {
		log.Fatal("empty db connection string")
	}

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	p.db = db
}

// Ping try to ping database
func (p *PostgresAdapter) Ping() bool {
	log.Println("trying to ping database...")
	err := p.db.Ping()
	if err != nil {
		log.Println("failed to ping db, err:", err)
		return false
	}

	return true
}

// GetDatabase returns *sql.DB instance
func (p *PostgresAdapter) GetDatabase() *sql.DB {
	return p.db
}
