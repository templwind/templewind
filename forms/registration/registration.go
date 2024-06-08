package registration

import (
	"github.com/a-h/templ"
	"github.com/templwind/templwind"
	"github.com/templwind/templwind/forms"
)

// Props defines the options for the AppBar component
type Props struct {
	ID     string
	Inputs []forms.Input
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

// DefaultProps provides the default options for the AppBar component
func defaultProps() *Props {
	return &Props{}
}

func WithID(id string) templwind.OptFunc[Props] {
	return func(opts *Props) {
		opts.ID = id
	}
}

func WithInputs(inputs []forms.Input) templwind.OptFunc[Props] {
	return func(opts *Props) {
		opts.Inputs = inputs
	}
}
