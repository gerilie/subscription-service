package validation

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

var ErrRequired = fmt.Errorf("%s required", ValidationPrefix)

// getErrorForRequiredTag returns a user-friendly error for required-related validation failures.
//
// It inspects the validation tag from the provided validator.FieldError
// and maps it to a corresponding error value.
//
// Supported tags:
//   - "required":      the field must be present
//   - "required_with": the field must be present when the specified field is present
//
// If the validation tag is not recognized, it returns nil.
func getErrorForRequiredTag(fe validator.FieldError) error {
	switch fe.Tag() {
	case "required":
		return ErrRequired
	case "required_with":
		return ErrRequiredWith(fe.Param())
	}

	return nil
}

func ErrRequiredWith(field string) error {
	return fmt.Errorf(
		"%s present when '%s' is present",
		ValidationPrefix,
		field,
	)
}
