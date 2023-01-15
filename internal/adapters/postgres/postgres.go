package postgres

import (
	"errors"
	"github.com/gocraft/dbr/v2"
	_ "github.com/lib/pq"
	"github.com/victoorraphael/coordinator/internal/adapters"
	"log"
	"os"
	"sync"
)

type Adapter struct {
	mu        sync.Mutex
	closed    bool
	resources chan *dbr.Session
	db        *dbr.Connection
}

// NewAdapter returns a new instance of DBPool
func NewAdapter(size uint) (adapters.DBPool, error) {
	if size <= 0 {
		return nil, errors.New("size too small")
	}

	return &Adapter{
		mu:        sync.Mutex{},
		closed:    false,
		resources: make(chan *dbr.Session, size),
		db:        connect(size),
	}, nil
}

// Acquire returns a new DB session to process queries
func (p *Adapter) Acquire() (*dbr.Session, error) {
	select {
	case r, ok := <-p.resources:
		if !ok {
			return nil, errors.New("pool closed")
		}
		return r, nil
	default:
		return p.factorySession(), nil
	}
}

// Release try to put connection back to the pool
// if not able to put back, connection is closed
func (p *Adapter) Release(resource *dbr.Session) {
	p.mu.Lock()
	defer p.mu.Unlock()

	if p.closed {
		_ = resource.Close()
		return
	}
	select {
	case p.resources <- resource:
		log.Println("resource in queue")
	default:
		_ = resource.Close()
	}
}

// Close finish the resources pool
func (p *Adapter) Close() {
	p.mu.Lock()
	defer p.mu.Unlock()

	if p.closed {
		return
	}

	p.closed = true

	close(p.resources)

	for r := range p.resources {
		_ = r.Close()
	}
}

// factorySession produces new sessions to be used into pool
func (p *Adapter) factorySession() *dbr.Session {
	return p.db.NewSession(nil)
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

// connect try to connect DB with environment variable
func connect(size uint) *dbr.Connection {
	connStr := os.Getenv("DB_URI")

	if connStr == "" {
		log.Fatal("empty db connection string")
	}

	db, err := dbr.Open("postgres", connStr, nil)
	if err != nil {
		log.Fatal(err)
	}

	db.SetMaxOpenConns(int(size))
	db.SetMaxIdleConns(int(size))
	return db
}
