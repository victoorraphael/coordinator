package entities

type AdaptersConfiguration func(a *Adapters) error

type Adapters struct {
	DB DB
}

type DB interface {
	Migrate()
	Ping() bool
}
