package card

import (
	"github.com/templwind/templwind"
	"github.com/templwind/templwind/components/indicator"

	"github.com/a-h/templ"
)

type Props struct {
	ID            string
	Class         string
	Title         string
	SubTitle      string
	Lead          string
	HeadIndicator *indicator.Props
	Components    []templ.Component
	Buttons       templ.Component
}

// New creates a new component
func New(prosps ...templwind.OptFunc[Props]) templ.Component {
	return templwind.New(defaultProps, tpl, prosps...)
}

// NewWithProps creates a new component with the given props
func NewWithProps(props *Props) templ.Component {
	return templwind.NewWithProps(tpl, props)
}

// WithProps builds the propsions with the given props
func WithProps(prosps ...templwind.OptFunc[Props]) *Props {
	return templwind.WithProps(defaultProps, prosps...)
}

func defaultProps() *Props {
	return &Props{
		Class: "bg-white border border-slate-200 rounded-lg shadow dark:bg-slate-800 dark:border-slate-700",
	}
}

func WithID(id string) templwind.OptFunc[Props] {
	return func(prosps *Props) {
		prosps.ID = id
	}
}

func WithClass(class string) templwind.OptFunc[Props] {
	return func(prosps *Props) {
		prosps.Class = class
	}
}

func WithTitle(title string) templwind.OptFunc[Props] {
	return func(prosps *Props) {
		prosps.Title = title
	}
}

func WithSubTitle(subTitle string) templwind.OptFunc[Props] {
	return func(prosps *Props) {
		prosps.SubTitle = subTitle
	}
}

func WithLead(lead string) templwind.OptFunc[Props] {
	return func(prosps *Props) {
		prosps.Lead = lead
	}
}

func WithHeadIndicator(headIndicator *indicator.Props) templwind.OptFunc[Props] {
	return func(prosps *Props) {
		prosps.HeadIndicator = headIndicator
	}
}

func WithComponents(components ...templ.Component) templwind.OptFunc[Props] {
	return func(prosps *Props) {
		prosps.Components = components
	}
}

func WithButtons(buttons templ.Component) templwind.OptFunc[Props] {
	return func(prosps *Props) {
		prosps.Buttons = buttons
	}
}
