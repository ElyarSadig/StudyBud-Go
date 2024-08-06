package errorHandler

import (
	"fmt"
)

var _ Handler = (*Error)(nil)

type Handler interface {
	New(httpCode int, message string, params ...any) *Error
	Error() string
	HTTPStatus() int
}

// NewError create error handler object
func NewError() (Handler, error) {
	e := new(Error)

	return e, nil
}

// New create error with custom message and params
func (e *Error) New(httpCode int, message string, params ...any) *Error {
	e.message = message
	e.params = params
	e.httpCode = httpCode

	return e
}

// Error show error message with appended parameters
func (e *Error) Error() string {
	if len(e.params) != 0 {
		return fmt.Sprintf(e.message, e.params)
	}
	return e.message
}

func (e *Error) HTTPStatus() int {
	return e.httpCode
}
