package indicator

import (
	"github.com/a-h/templ"
	"github.com/templwind/templwind"
)

type Props struct {
	IsUp  bool
	Value string
}

// New creates a new component
func New(props ...templwind.OptFunc[Props]) templ.Component {
	return templwind.New(defaultProps, tpl, props...)
}

// NewWithOpt creates a new component with the given props
func NewWithProps(props *Props) templ.Component {
	return templwind.NewWithProps(tpl, props)
}

// WithProps builds the propsions with the given props
func WithProps(props ...templwind.OptFunc[Props]) *Props {
	return templwind.WithProps(defaultProps, props...)
}

func defaultProps() *Props {
	return &Props{}
}

func WithIsUp(props *Props) {
	props.IsUp = true
}

func WithValue(value string) templwind.OptFunc[Props] {
	return func(props *Props) {
		props.Value = value
	}
}
