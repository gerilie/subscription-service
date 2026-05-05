package validation

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

// getErrorForTag returns a user-friendly error for a validation failure.
//
// It attempts to resolve the error by delegating to specific handlers
// based on the validation tag (e.g., required, string, number, datetime).
// The first non-nil error returned by these handlers is used.
//
// If no handler matches the tag, it returns a default error message
// containing the validation tag and field name.
func getErrorForTag(fe validator.FieldError) error {
	err := getErrorForRequiredTag(fe)
	if err != nil {
		return err
	}

	err = getErrorForStringTag(fe)
	if err != nil {
		return err
	}

	err = getErrorForNumberTag(fe)
	if err != nil {
		return err
	}

	err = getErrorForDatetimeTag(fe)
	if err != nil {
		return err
	}

	return fmt.Errorf(
		"validation failed for the '%s' rule on field '%s'",
		fe.Tag(),
		fe.Field(),
	)
}
