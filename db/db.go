package db

import (
	"math"
	"strings"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/xo/dburl"
)

// PersistentSQLx contains the persistent database connection
type PersistentSQLx struct {
	db   *sqlx.DB
	dsn  string
	opts *DBConfig
}

// OptFunc defines the signature for an option function
type OptFunc[T any] func(*T)

// Connect creates a new PersistentSQLx instance
func MustConnect(opts ...OptFunc[DBConfig]) *PersistentSQLx {
	options := WithOptions(defaultOptions, opts...)
	return NewWithOptions(options)
}

// NewWithOptions creates a new PersistentSQLx instance with the given options
func NewWithOptions(opts *DBConfig) *PersistentSQLx {
	dsn := strings.ReplaceAll(opts.DSN, "\"", "")
	db, err := connect(opts)
	if err != nil {
		panic("Failed to connect to database: " + err.Error())
	}

	psqlx := &PersistentSQLx{
		db:   db,
		dsn:  dsn,
		opts: opts,
	}

	// Start a go-routine to continuously check connection health
	go psqlx.ensureConnection()

	return psqlx
}

// WithOptions builds the options with the given opt
func WithOptions(defaultOpts func() *DBConfig, opts ...OptFunc[DBConfig]) *DBConfig {
	p := defaultOpts()
	for _, opt := range opts {
		opt(p)
	}
	return p
}

// defaultOptions returns the default options for PersistentSQLx
func defaultOptions() *DBConfig {
	return &DBConfig{
		EnableWALMode: false,
	}
}

// WithDSN sets the DSN
func WithDSN(dsn string) OptFunc[DBConfig] {
	return func(p *DBConfig) {
		p.DSN = dsn
	}
}

// WithEnableWALMode sets the WAL mode
func WithEnableWALMode(enable bool) OptFunc[DBConfig] {
	return func(p *DBConfig) {
		p.EnableWALMode = enable
	}
}

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

// GetDB returns the database connection
func (psqlx *PersistentSQLx) GetDB() *sqlx.DB {
	return psqlx.db
}

// reconnect attempts to re-establish the database connection with exponential backoff
func (psqlx *PersistentSQLx) reconnect() error {
	// Exponential backoff parameters
	const maxBackoff = 5 * time.Minute
	baseDelay := 500 * time.Millisecond

	for attempts := 0; ; attempts++ {
		db, err := connect(psqlx.opts)
		if err == nil {
			psqlx.db = db
			return nil
		}

		if attempts > 0 {
			// Exponential backoff calculation
			backoff := time.Duration(math.Pow(2, float64(attempts))) * baseDelay
			if backoff > maxBackoff {
				backoff = maxBackoff
			}

			time.Sleep(backoff)
		}
	}
}

// ensureConnection continuously checks the health of the database connection
func (psqlx *PersistentSQLx) ensureConnection() {
	for {
		// Implement a simple ping check
		if err := psqlx.db.Ping(); err != nil {
			// If ping fails, try to reconnect
			psqlx.reconnect()
		}

		// Sleep for some time before the next health check
		time.Sleep(1 * time.Minute)
	}
}
