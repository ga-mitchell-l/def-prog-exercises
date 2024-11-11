package safesql

import (
	"context"
	"database/sql"

	"github.com/ga-mitchell-l/def-prog-exercises/authentication"
	"github.com/ga-mitchell-l/def-prog-exercises/safesql/internal/raw"
)

type compileTimeConstant string

type TrustedSQL struct {
	s string
}

// stops you passing run time strings
func New(text compileTimeConstant) TrustedSQL {
	return TrustedSQL{string(text)}
}

func init() {
	raw.TrustedSQLCtor = func(unsafe string) TrustedSQL {
		return TrustedSQL{unsafe}
	}
}

type DB struct {
	db *sql.DB
}

func (db *DB) QueryContext(ctx context.Context,
	query TrustedSQL, args ...any) (*Rows, error) {
	authentication.Must(ctx)
	r, err := db.db.QueryContext(ctx, query.s, args...)
	return r, err
}

func (db *DB) ExecContext(ctx context.Context,
	query TrustedSQL, args ...any) (Result, error) {
	authentication.Must(ctx)
	return db.db.ExecContext(ctx, query.s, args...)
}

func (db *DB) Open(driverName string, dataSourceName string) (*sql.DB, error) {
	return sql.Open(driverName, dataSourceName)
}

type (
	Result = sql.Result
	Rows   = sql.Rows
)

// Also wrap: sql.db.ExecContext and sql.Open
// Also alias: sql.Result
