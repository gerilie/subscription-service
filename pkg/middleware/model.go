package middleware

import "errors"

// ErrMiddleware is the base error for all middleware errors.
var ErrMiddleware = errors.New("middleware error")
