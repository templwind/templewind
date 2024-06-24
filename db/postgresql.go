//go:build postgresql
// +build postgresql

package db

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/xo/dburl"
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

	return dbConn, nil
}
