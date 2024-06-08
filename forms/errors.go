package forms

import (
	"errors"
	"fmt"
)

// ErrUnsupportedMediaType is returned when the request content type is not supported.
var ErrUnsupportedMediaType = errors.New("unsupported media type")

// HTTPError represents an error that occurred while handling a request.
type HTTPError struct {
	Code     int         `json:"code"`
	Message  interface{} `json:"message"`
	Internal error       `json:"-"` // Stores the internal error, if any.
}

// NewHTTPError creates a new HTTPError instance.
func NewHTTPError(code int, message interface{}) *HTTPError {
	return &HTTPError{
		Code:    code,
		Message: message,
	}
}

// Error makes it compatible with `error` interface.
func (e *HTTPError) Error() string {
	return fmt.Sprintf("code=%d, message=%v", e.Code, e.Message)
}

// Unwrap provides compatibility for Go 1.13 error chains.
func (e *HTTPError) Unwrap() error {
	return e.Internal
}

// SetInternal sets the internal error and returns the HTTPError.
func (e *HTTPError) SetInternal(err error) *HTTPError {
	e.Internal = err
	return e
}
