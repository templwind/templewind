//go:build mssql
// +build mssql

package db

import (
	_ "github.com/denisenkom/go-mssqldb"
	"github.com/jmoiron/sqlx"
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
