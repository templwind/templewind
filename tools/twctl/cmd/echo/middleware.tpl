package middleware

import (
	{{.ImportPackages}}
)

type {{.name}} struct {
}

func New{{.name}}() *{{.name}} {
	return &{{.name}}{}
}

func (m *{{.name}})Handle(next echo.HandlerFunc) echo.HandlerFunc {
	return func(e echo.Context) error {
		// TODO generate middleware implement function, delete after code implementation

		// Passthrough to next handler if need
		return next(e)
	}
}
