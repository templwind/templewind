package sidenav

import (
	"github.com/a-h/templ"
	"github.com/templwind/templwind"
	"github.com/templwind/templwind/components/link"
)

type Props struct {
	ID             string
	ContainerClass string
	Submenu        []*link.Props
}

// New creates a new component
func New(props ...templwind.OptFunc[Props]) templ.Component {
	return templwind.New(defaultProps, tpl, props...)
}

// NewWithProps creates a new component with the given prop
func NewWithProps(props *Props) templ.Component {
	return templwind.NewWithProps(tpl, props)
}

// WithProps builds the propions with the given prop
func WithProps(props ...templwind.OptFunc[Props]) *Props {
	return templwind.WithProps(defaultProps, props...)
}

// DefaultProps provides the default propions for the AppBar component
func defaultProps() *Props {
	return &Props{
		ID:      "sidenav",
		Submenu: []*link.Props{},
	}
}

func WithID(id string) templwind.OptFunc[Props] {
	return func(props *Props) {
		props.ID = id
	}
}

func WithContainerClass(class string) templwind.OptFunc[Props] {
	return func(props *Props) {
		props.ContainerClass = class
	}
}

func WithSubmenu(submenu ...*link.Props) templwind.OptFunc[Props] {
	return func(props *Props) {
		props.Submenu = append(props.Submenu, submenu...)
	}
}
