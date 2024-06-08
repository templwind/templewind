package forms

import "net/http"

var defaultBinder = DefaultBinder{}

// Bind is a convenience function that uses the default binder to bind the request data to the given struct.
func Bind(req *http.Request, i interface{}) error {
	return defaultBinder.Bind(i, req)
}
