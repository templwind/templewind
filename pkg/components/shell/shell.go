package shell

import (
	"github.com/a-h/templ"
	"github.com/templwind/templwind"
)

// Props for the shell component
type Props struct {
	ID           string
	Header       templ.Component
	SidebarLeft  templ.Component
	SidebarRight templ.Component
	PageHeader   templ.Component
	PageFooter   templ.Component
	Footer       templ.Component
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
		ID: "shell",
	}
}

func WithID(id string) templwind.OptFunc[Props] {
	return func(p *Props) {
		p.ID = id
	}
}

func WithHeader(c templ.Component) templwind.OptFunc[Props] {
	return func(p *Props) {
		p.Header = c
	}
}

func WithSidebarLeft(c templ.Component) templwind.OptFunc[Props] {
	return func(p *Props) {
		p.SidebarLeft = c
	}
}

func WithSidebarRight(c templ.Component) templwind.OptFunc[Props] {
	return func(p *Props) {
		p.SidebarRight = c
	}
}

func WithPageHeader(c templ.Component) templwind.OptFunc[Props] {
	return func(p *Props) {
		p.PageHeader = c
	}
}

func WithPageFooter(c templ.Component) templwind.OptFunc[Props] {
	return func(p *Props) {
		p.PageFooter = c
	}
}

func WithFooter(c templ.Component) templwind.OptFunc[Props] {
	return func(p *Props) {
		p.Footer = c
	}
}
