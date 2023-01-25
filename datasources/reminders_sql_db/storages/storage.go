package storages

import (
	"context"
	"database/sql"
)

type IReminderSQLDBStorage interface {
	Master() *sql.DB
	Slave() *sql.DB
	BeginTx(ctx context.Context) (*sql.Tx, error)
}
