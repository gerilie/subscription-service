package validation

import "errors"

// Resp maps field names to user-friendly validation error messages.
//
// Example:
//
//	{
//		"email": "must be a valid email address",
//		"age":   "must be greater than or equal to 18",
//	}
type Resp map[string]string

// ErrValidation is the base error for all validation failures.
var ErrValidation = errors.New("validation failed")
