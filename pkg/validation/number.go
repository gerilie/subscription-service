package validation

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

// getErrorMessageForNumberTag returns a user-friendly error message
// for numeric validation errors.
//
// It inspects the validation tag from the provided validator.FieldError
// and maps it to a corresponding message.
//
// Supported tags:
//   - "min": the field must be greater than or equal to the specified value
//   - "max": the field must be less than or equal to the specified value
//   - "gt":  the field must be greater than the specified value
//   - "gte": the field must be greater than or equal to the specified value
//
// If the validation tag is not recognized, getErrorMessageForNumberTag returns an empty string.
func getErrorMessageForNumberTag(fe validator.FieldError) string {
	switch fe.Tag() {
	case "min":
		return fmt.Sprintf(
			"The field must be greater than or equal to %s",
			fe.Param(),
		)
	case "max":
		return fmt.Sprintf(
			"The field must be less than or equal to %s",
			fe.Param(),
		)
	case "gt":
		return fmt.Sprintf(
			"The field must be greater than %s",
			fe.Param(),
		)
	case "gte":
		return fmt.Sprintf(
			"The field must be greater than or equal to %s",
			fe.Param(),
		)
	}

	return ""
}
