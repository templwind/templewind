package appbar

import (
	"github.com/a-h/templ"
	"github.com/templwind/templwind"
)

type Class interface {
	Border() string
	BorderTop() string
	BorderBottom() string
	BorderLeft() string
	BorderRight() string
}

// Props defines the options for the AppBar component
type Props struct {
	AppBarClasses   string
	Lead            templ.Component
	LeadClasses     string
	Trail           templ.Component
	TrailClasses    string
	Headline        templ.Component
	HeadlineClasses string
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
	return &Props{
		AppBarClasses:   "app-bar flex flex-col bg-surface-100-800-token space-y-4 p-4 w-full",
		LeadClasses:     "app-bar-slot-lead flex-none flex justify-between items-center",
		TrailClasses:    "app-bar-slot-trail flex-none flex items-center space-x-4",
		HeadlineClasses: "app-bar-row-headline",
	}
}

func WithAppBarClasses(appBarClasses string) templwind.OptFunc[Props] {
	return func(opts *Props) {
		opts.AppBarClasses = appBarClasses
	}
}

func WithLead(lead templ.Component) templwind.OptFunc[Props] {
	return func(opts *Props) {
		opts.Lead = lead
	}
}

func WithLeadClasses(leadClasses string) templwind.OptFunc[Props] {
	return func(opts *Props) {
		opts.LeadClasses = leadClasses
	}
}

func WithTrail(trail templ.Component) templwind.OptFunc[Props] {
	return func(opts *Props) {
		opts.Trail = trail
	}
}

func WithTrailClasses(trailClasses string) templwind.OptFunc[Props] {
	return func(opts *Props) {
		opts.TrailClasses = trailClasses
	}
}

func WithHeadline(headline templ.Component) templwind.OptFunc[Props] {
	return func(opts *Props) {
		opts.Headline = headline
	}
}

func WithHeadlineClasses(headlineClasses string) templwind.OptFunc[Props] {
	return func(opts *Props) {
		opts.HeadlineClasses = headlineClasses
	}
}
