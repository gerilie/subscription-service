package validation

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

// getErrorForTag maps a validator.FieldError
// to a user-friendly validation error.
//
// Validation tags are resolved by specialized handlers
// in the following order:
//   - required validation tags
//   - string validation tags
//   - numeric validation tags
//   - datetime validation tags
//
// The first matching non-nil error is returned.
//
// If the validation tag is unsupported, a generic fallback
// error describing the failed rule and field is returned.
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
		"%w: rule '%s', field '%s'",
		ErrValidation,
		fe.Tag(),
		fe.Field(),
	)
}
