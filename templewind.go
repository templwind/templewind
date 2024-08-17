package templwind

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer/html"
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
	fmt.Printf("prop: %v\n", prop)
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
	md := goldmark.New(
		goldmark.WithExtensions(
			extension.GFM,           // GitHub Flavored Markdown (tables, strikethrough, etc.)
			extension.Linkify,       // Automatically turns URLs into links
			extension.Strikethrough, // Strikethrough support
		),
		goldmark.WithParserOptions(
			parser.WithAutoHeadingID(), // Automatically generates heading IDs
		),
		goldmark.WithRendererOptions(
			html.WithUnsafe(), // Allows rendering of raw HTML in Markdown
		),
	)

	var buf bytes.Buffer
	if err := md.Convert([]byte(markdown), &buf); err != nil {
		log.Printf("failed to convert markdown to HTML: %v", err)
	}
	return Unsafe(buf.String())
}
