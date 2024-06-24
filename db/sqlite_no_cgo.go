//go:build sqlite_no_cgo
// +build sqlite_no_cgo

package db

import (
	_ "modernc.org/sqlite"
)
