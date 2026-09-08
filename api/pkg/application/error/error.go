package error

import (
	"errors"
	"fmt"
)

type appError struct {
	err        error
	code       string
	message    string
	statusCode int
	debugData  map[string]any
}

func NewAppError(err error, formatValues ...any) *appError {
	if err == nil {
		return nil
	}
	appErr := &appError{
		err: err,
	}
	for definedErr, defined := range definedErrors {
		if errors.Is(err, definedErr) {
			appErr.code = definedErr.Error()
			appErr.message = fmt.Sprintf(defined.message, formatValues...)
			appErr.statusCode = defined.statusCode
			return appErr
		}
	}
	appErr.code = ErrInternalServer.Error()
	appErr.message = fmt.Sprintf(definedErrors[ErrInternalServer].message, formatValues...)
	appErr.statusCode = definedErrors[ErrInternalServer].statusCode
	return appErr
}

func (e *appError) Error() string {
	return e.err.Error()
}

func (e *appError) Code() string {
	return e.code
}

func (e *appError) Message() string {
	return e.message
}

func (e *appError) StatusCode() int {
	return e.statusCode
}
