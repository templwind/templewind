package templwind

import (
	"bytes"
	"context"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
	"github.com/yuin/goldmark"
)

// Component interface with generic methods
type Component[T any] interface {
	New(props ...OptFunc[T]) templ.Component
	NewWithOpt(prop *T) templ.Component
	WithProps(props ...OptFunc[T]) *T
}

// OptFunc is a generic function type for props
type OptFunc[T any] func(*T)

// New creates a new templ.Component with the given props
func New[T any](defaultProps func() *T, tpl func(*T) templ.Component, props ...OptFunc[T]) templ.Component {
	prop := WithProps(defaultProps, props...)
	return tpl(prop)
}

// NewWithProps creates a new templ.Component with the given prop
func NewWithProps[T any](tpl func(*T) templ.Component, props *T) templ.Component {
	return tpl(props)
}

// WithProps constructs the props with the given prop functions
func WithProps[T any](defaultProps func() *T, props ...OptFunc[T]) *T {
	defaults := defaultProps()
	for _, propFn := range props {
		propFn(defaults)
	}
	return defaults
}

func Render(ctx echo.Context, status int, t templ.Component) error {
	if !ctx.Response().Committed {
		ctx.Response().WriteHeader(status)
	}

	err := t.Render(context.Background(), ctx.Response().Writer)
	if err != nil {
		log.Println(err)
		return ctx.String(http.StatusInternalServerError, "failed to render response template")
	}

	return nil
}

func ComponentToString(c templ.Component) (string, error) {
	var sb strings.Builder
	err := c.Render(context.Background(), &sb)
	if err != nil {
		return "", err
	}
	return sb.String(), nil
}

func Unsafe(html string) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) (err error) {
		_, err = io.WriteString(w, html)
		return
	})
}

func Markdown(markdown string) templ.Component {
	var buf bytes.Buffer
	if err := goldmark.Convert([]byte(markdown), &buf); err != nil {
		log.Printf("failed to convert markdown to HTML: %v", err)
	}
	return Unsafe(buf.String())
}
