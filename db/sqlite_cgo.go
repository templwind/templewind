//go:build cgo && sqlite

package db

import (
	_ "github.com/mattn/go-sqlite3"
)
