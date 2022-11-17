package adapters

import "database/sql"

type DBAdapter interface {
	Ping() bool
	GetDatabase() *sql.DB
}
