package validation

import (
	"github.com/go-playground/validator/v10"
)

// formatErrorsByName converts validator.ValidationErrors
// into a map of field names to user-friendly error messages.
//
// Each validation error is mapped using getErrorForTag.
// Errors with unsupported validation tags are skipped.
//
// Example result:
//
//	map[string]string{
//		"Email": "must be a valid email address",
//		"Age":   "must be greater than or equal to 18",
//	}
func formatErrorsByName(ve validator.ValidationErrors) Resp {
	errors := make(Resp)

	for _, fe := range ve {
		name := fe.Field()

		err := getErrorForTag(fe)
		if err == nil {
			continue
		}

		errors[name] = err.Error()
	}

	return errors
}
