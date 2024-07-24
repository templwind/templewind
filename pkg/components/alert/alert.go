package alert

import (
	"fmt"

	"github.com/a-h/templ"
	"github.com/rs/xid"
	"github.com/templwind/soul/util"
	"github.com/templwind/templwind"
)

type AlertType string

const (
	Info    AlertType = "info"
	Success AlertType = "success"
	Warning AlertType = "warning"
	Error   AlertType = "error"
)

func (t AlertType) IsInfo() bool {
	return t == Info
}

func (t AlertType) IsSuccess() bool {
	return t == Success
}

func (t AlertType) IsWarning() bool {
	return t == Warning
}

func (t AlertType) IsError() bool {
	return t == Error
}

// Props for the alert component
type Props struct {
	ID           string
	Type         AlertType // alert type: info, success, warning, error
	Message      string    // alert message
	HideDuration int       // duration in milliseconds to hide the alert automatically
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
		ID:           util.ToCamel(fmt.Sprintf("alert-%s", xid.New().String())),
		Type:         Info,
		Message:      "This is an alert",
		HideDuration: 3000, // default to 3 seconds
	}
}

func WithID(id string) templwind.OptFunc[Props] {
	return func(p *Props) {
		p.ID = id
	}
}

func WithType(t AlertType) templwind.OptFunc[Props] {
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
