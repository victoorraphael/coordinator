package postgres

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"os"
)

type Adapter struct {
	db *sql.DB
}

func NewPostgresAdapter() *Adapter {
	p := &Adapter{}
	//TODO include pool of resources instead of connect every time
	p.connect()
	return p
}

// connect try to connect DB with environment variable
func (p *Adapter) connect() {
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
func (p *Adapter) Ping() bool {
	log.Println("trying to ping database...")
	err := p.db.Ping()
	if err != nil {
		log.Println("failed to ping db, err:", err)
		return false
	}

	return true
}

// GetDatabase returns *sql.DB instance
func (p *Adapter) GetDatabase() *sql.DB {
	return p.db
}
