package validation

import (
	"github.com/go-playground/validator/v10"
)

// formatErrorsByName converts validator.ValidationErrors into a Resp map.
//
// It iterates over validation errors, maps each struct field name to a
// user-friendly error message produced by getErrorForTag, and skips
// any errors for which getErrorForTag returns nil.
//
// The resulting map uses field names as keys and error messages as values.
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
