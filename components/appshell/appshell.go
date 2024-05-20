package appshell

import (
	"github.com/a-h/templ"
	"github.com/templwind/templwind"
)

// Options struct defines the configuration for the AppShell component.
type Props struct {
	Header              templ.Component
	HeaderClasses       string
	SidebarLeft         templ.Component
	SidebarLeftClasses  string
	SidebarRight        templ.Component
	SidebarRightClasses string
	PageHeader          templ.Component
	PageHeaderClasses   string
	PageFooter          templ.Component
	PageFooterClasses   string
	Footer              templ.Component
	FooterClasses       string
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

// Default propions for the AppShell component.
func defaultProps() *Props {
	return &Props{
		Header:       nil,
		SidebarLeft:  nil,
		SidebarRight: nil,
	}
}

func WithHeader(header templ.Component) templwind.OptFunc[Props] {
	return func(props *Props) {
		props.Header = header
	}
}

func WithHeaderClasses(headerClasses string) templwind.OptFunc[Props] {
	return func(props *Props) {
		props.HeaderClasses = headerClasses
	}
}

func WithSidebarLeft(sidebarLeft templ.Component) templwind.OptFunc[Props] {
	return func(props *Props) {
		props.SidebarLeft = sidebarLeft
	}
}

func WithSidebarLeftClasses(sidebarLeftClasses string) templwind.OptFunc[Props] {
	return func(props *Props) {
		props.SidebarLeftClasses = sidebarLeftClasses
	}
}

func WithSidebarRight(sidebarRight templ.Component) templwind.OptFunc[Props] {
	return func(props *Props) {
		props.SidebarRight = sidebarRight
	}
}

func WithSidebarRightClasses(sidebarRightClasses string) templwind.OptFunc[Props] {
	return func(props *Props) {
		props.SidebarRightClasses = sidebarRightClasses
	}
}

func WithPageHeader(pageHeader templ.Component) templwind.OptFunc[Props] {
	return func(props *Props) {
		props.PageHeader = pageHeader
	}
}

func WithPageHeaderClasses(pageHeaderClasses string) templwind.OptFunc[Props] {
	return func(props *Props) {
		props.PageHeaderClasses = pageHeaderClasses
	}
}

func WithPageFooter(pageFooter templ.Component) templwind.OptFunc[Props] {
	return func(props *Props) {
		props.PageFooter = pageFooter
	}
}

func WithPageFooterClasses(pageFooterClasses string) templwind.OptFunc[Props] {
	return func(props *Props) {
		props.PageFooterClasses = pageFooterClasses
	}
}

func WithFooter(footer templ.Component) templwind.OptFunc[Props] {
	return func(props *Props) {
		props.Footer = footer
	}
}

func WithFooterClasses(footerClasses string) templwind.OptFunc[Props] {
	return func(props *Props) {
		props.FooterClasses = footerClasses
	}
}
