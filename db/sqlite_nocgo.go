//go:build !cgo

package db

import (
	"github.com/jmoiron/sqlx"
	"github.com/xo/dburl"
	_ "modernc.org/sqlite"
)

// connect establishes a new database connection
func connect(opts *DBConfig) (*sqlx.DB, error) {
	u, err := dburl.Parse(opts.DSN)
	if err != nil {
		return nil, err
	}

	dbConn, err := sqlx.Open(u.Driver, u.DSN)
	if err != nil {
		return nil, err
	}

	// Enable WAL mode if SQLite and requested
	if u.Driver == "sqlite3" && opts.EnableWALMode {
		_, err = dbConn.Exec("PRAGMA journal_mode = WAL;")
		if err != nil {
			return nil, err
		}
	}

	return dbConn, nil
}
