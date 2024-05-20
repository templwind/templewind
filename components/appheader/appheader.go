package appheader

import (
	"fmt"

	"github.com/a-h/templ"
	"github.com/google/uuid"
	"github.com/templwind/templwind"
	"github.com/templwind/templwind/components/link"
	"github.com/templwind/templwind/utils"
)

type Props struct {
	ID           string
	HideOnMobile bool
	LinkProps    *link.Props
	Title        string
}

// New creates a new component
func New(opts ...templwind.OptFunc[Props]) templ.Component {
	return templwind.New(defaultProps, tpl, opts...)
}

// NewWithProps creates a new component with the given opt
func NewWithProps(opt *Props) templ.Component {
	return templwind.NewWithProps(tpl, opt)
}

// WithProps builds the options with the given opt
func WithProps(opts ...templwind.OptFunc[Props]) *Props {
	return templwind.WithProps(defaultProps, opts...)
}

func defaultProps() *Props {
	return &Props{
		ID:           utils.ToCamel(fmt.Sprintf("appHeader-%s", uuid.New().String())),
		HideOnMobile: false,
		Title:        "App Header",
	}
}

func WithID(id string) templwind.OptFunc[Props] {
	return func(o *Props) {
		o.ID = id
	}
}

func WithHideOnMobile(hide bool) templwind.OptFunc[Props] {
	return func(o *Props) {
		o.HideOnMobile = hide
	}
}

func WithLinkProps(linkProps ...templwind.OptFunc[link.Props]) templwind.OptFunc[Props] {
	return func(o *Props) {
		o.LinkProps = link.WithProps(linkProps...)
	}
}

func WithTitle(title string) templwind.OptFunc[Props] {
	return func(o *Props) {
		o.Title = title
	}
}
