package navbar

import (
	"github.com/a-h/templ"
	"github.com/templwind/templwind"
)

type MenuItem struct {
	Icon  string
	Title string
	Link  string
}

type Props struct {
	ID        string
	BrandName string
	BrandLogo string
	BrandLink string
	MenuItems []MenuItem
}

// New creates a new component
func New(props ...templwind.OptFunc[Props]) templ.Component {
	return templwind.New(defaultProps, tpl, props...)
}

// NewWithProps creates a new component with the given prosp
func NewWithProps(props *Props) templ.Component {
	return templwind.NewWithProps(tpl, props)
}

// WithProps builds the prospions with the given prosp
func WithProps(props ...templwind.OptFunc[Props]) *Props {
	return templwind.WithProps(defaultProps, props...)
}

func defaultProps() *Props {
	return &Props{}
}

func WithID(id string) templwind.OptFunc[Props] {
	return func(props *Props) {
		props.ID = id
	}
}

func WithBrandName(name string) templwind.OptFunc[Props] {
	return func(o *Props) {
		o.BrandName = name
	}
}

func WithBrandLogo(logo string) templwind.OptFunc[Props] {
	return func(o *Props) {
		o.BrandLogo = logo
	}
}

func WithBrandURL(link string) templwind.OptFunc[Props] {
	return func(o *Props) {
		o.BrandLink = link
	}
}

func WithMenu(items []MenuItem) templwind.OptFunc[Props] {
	return func(o *Props) {
		o.MenuItems = items
	}
}
