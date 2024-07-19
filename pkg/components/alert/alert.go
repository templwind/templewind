package alert

import (
	"github.com/a-h/templ"
	"github.com/templwind/templwind"
)

// Props for the alert component
type Props struct {
	Type         string // alert type: info, success, warning, error
	Message      string // alert message
	HideDuration int    // duration in milliseconds to hide the alert automatically
}

// New creates a new component
func New(props ...templwind.OptFunc[Props]) templ.Component {
	return templwind.New(defaultProps, tpl, props...)
}

// NewWithProps creates a new component with the given props
func NewWithProps(props *Props) templ.Component {
	return templwind.NewWithProps(tpl, props)
}

// WithProps builds the props with the given options
func WithProps(props ...templwind.OptFunc[Props]) *Props {
	return templwind.WithProps(defaultProps, props...)
}

func defaultProps() *Props {
	return &Props{
		Type:         "info",
		Message:      "This is an alert",
		HideDuration: 3000, // default to 3 seconds
	}
}

func WithType(t string) templwind.OptFunc[Props] {
	return func(p *Props) {
		p.Type = t
	}
}

func WithMessage(m string) templwind.OptFunc[Props] {
	return func(p *Props) {
		p.Message = m
	}
}

func WithHideDuration(d int) templwind.OptFunc[Props] {
	return func(p *Props) {
		p.HideDuration = d
	}
}
