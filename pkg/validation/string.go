package validation

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

var (
	// ErrEmail indicates that a field value is not a valid email address.
	ErrEmail = fmt.Errorf(
		"%w: must be a valid email address",
		ErrValidation,
	)

	// ErrUUID4 indicates that a field value is not a valid UUID (version 4).
	ErrUUID4 = fmt.Errorf(
		"%w: must be a valid UUID v4",
		ErrValidation,
	)
)

// getErrorForStringTag maps validator string-related tags
// to user-friendly validation errors.
//
// Supported validator tags:
//   - "email":  field value must be a valid email address
//   - "uuid4":  field value must be a valid UUID (version 4)
//
// Returns nil if the validation tag is unsupported.
func getErrorForStringTag(fe validator.FieldError) error {
	switch fe.Tag() {
	case "email":
		return ErrEmail
	case "uuid4":
		return ErrUUID4
	}

	return nil
}
