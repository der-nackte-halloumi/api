package sql_datastore

import (
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"
)

type Executor interface {
	ExecContext(ctx context.Context, sql string, args ...interface{}) (sql.Result, error)
	QueryRowxContext(ctx context.Context, sql string, args ...interface{}) *sqlx.Row
	QueryxContext(ctx context.Context, sql string, args ...interface{}) (*sqlx.Rows, error)
}

type Datastore interface {
	Beginx() (*sqlx.Tx, error)
	Close() error
	Executor
}