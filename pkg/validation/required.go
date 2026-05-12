package validation

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

var (
	// ErrRequired indicates that a required field is missing.
	ErrRequired = fmt.Errorf(
		"%w: field is required",
		ErrValidation,
	)

	// ErrRequiredWith indicates that a field is required
	// when another field is present.
	ErrRequiredWith = fmt.Errorf(
		"%w: field is required when another field is present",
		ErrValidation,
	)
)

// getErrorForRequiredTag maps validator required-related tags
// to user-friendly validation errors.
//
// Supported validator tags:
//   - "required": field must be present and non-empty
//   - "required_with": field must be present when another field is present
//
// Returns nil if the validation tag is unsupported.
func getErrorForRequiredTag(fe validator.FieldError) error {
	switch fe.Tag() {
	case "required":
		return ErrRequired
	case "required_with":
		return RequiredWithError(fe.Param())
	}

	return nil
}

// RequiredWithError returns ErrRequiredWith enriched
// with the dependent field name.
func RequiredWithError(field string) error {
	return fmt.Errorf(
		"%w: required with '%s'",
		ErrRequiredWith,
		field,
	)
}
