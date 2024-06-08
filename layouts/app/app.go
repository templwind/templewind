package app

import (
	"net/http"

	"github.com/a-h/templ"
	"github.com/templwind/templwind"
	"github.com/templwind/templwind/components/appshell"
	"github.com/templwind/templwind/layouts/base"
)

// Options struct defines the configuration for the AppShell component.
type Props struct {
	base.Props
	AppShellProps []templwind.OptFunc[appshell.Props]
}

// Layout creates a new component
func Layout(props ...templwind.OptFunc[Props]) templ.Component {
	return templwind.New(defaultProps, tpl, props...)
}

// LayoutWithProps creates a new component with the given prop
func LayoutWithProps(props *Props) templ.Component {
	return templwind.NewWithProps(tpl, props)
}

// WithProps builds the props with the given prop
func WithProps(props ...templwind.OptFunc[Props]) *Props {
	return templwind.WithProps(defaultProps, props...)
}

// Default props for the AppShell component.
func defaultProps() *Props {
	return &Props{
		Props: base.Props{
			Title:         "TemplWind",
			Meta:          nil,
			Favicon:       "",
			Stylesheets:   []string{},
			HeadScripts:   []string{},
			FooterScripts: []string{},
			BodyCss:       "",
		},
	}
}

func WithHttpRequest(r *http.Request) templwind.OptFunc[Props] {
	return func(o *Props) {
		o.HttpRequest = r
	}
}

func WithHttpResponse(w http.ResponseWriter) templwind.OptFunc[Props] {
	return func(o *Props) {
		o.HttpResponse = w
	}
}

func WithBaseProps(baseProps ...templwind.OptFunc[base.Props]) templwind.OptFunc[Props] {
	return func(o *Props) {
		for _, bp := range baseProps {
			bp(&o.Props)
		}
	}
}

func WithAppShellProps(appShellProps ...templwind.OptFunc[appshell.Props]) templwind.OptFunc[Props] {
	return func(o *Props) {
		o.AppShellProps = appShellProps
	}
}

func WithTitle(title string) templwind.OptFunc[Props] {
	return func(o *Props) {
		o.Title = title
	}
}

func WithMeta(meta templ.Component) templwind.OptFunc[Props] {
	return func(o *Props) {
		o.Meta = meta
	}
}

func WithFavicon(favicon string) templwind.OptFunc[Props] {
	return func(o *Props) {
		o.Favicon = favicon
	}
}

func WithStylesheets(stylesheets ...string) templwind.OptFunc[Props] {
	return func(o *Props) {
		o.Stylesheets = stylesheets
	}
}

func WithHeadScripts(headScripts ...string) templwind.OptFunc[Props] {
	return func(o *Props) {
		o.HeadScripts = headScripts
	}
}

func WithFooterScripts(footerScripts ...string) templwind.OptFunc[Props] {
	return func(o *Props) {
		o.FooterScripts = footerScripts
	}
}

func WithBodyCss(bodyCss string) templwind.OptFunc[Props] {
	return func(o *Props) {
		o.BodyCss = bodyCss
	}
}
