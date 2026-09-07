package error

import (
	"errors"
)

var (
	ErrBadRequest         = errors.New("ERR_BAD_REQUEST")
	ErrUnauthorized       = errors.New("ERR_UNAUTHORIZED")
	ErrBearerTokenExpired = errors.New("ERR_BEARER_TOKEN_EXPIRED")
	ErrForbidden          = errors.New("ERR_FORBIDDEN")
	ErrNotFound           = errors.New("ERR_NOT_FOUND")
	ErrMethodNotAllowed   = errors.New("ERR_METHOD_NOT_ALLOWED")
	ErrInternalServer     = errors.New("ERR_INTERNAL_SERVER")
)
