package svc

import (
	{{.imports}}
)

type ServiceContext struct {
	Config {{.config}}
	DB         *sqlx.DB
	{{.middleware}}
}

func NewServiceContext(c {{.config}}) *ServiceContext {
	return &ServiceContext{
		Config: c,
		DB: db.MustConnect(
			db.WithDSN(c.DSN),
			db.WithEnableWALMode(true), // Enable WAL mode if needed
		).GetDB(),
		{{.middlewareAssignment}}
	}
}
