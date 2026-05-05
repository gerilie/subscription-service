package validation

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

var (
	ErrEmail = fmt.Errorf("%s a valid email address", ValidationPrefix)
	ErrUUID4 = fmt.Errorf("%s a valid UUID (version 4)", ValidationPrefix)
)

// getErrorForStringTag returns a user-friendly error for string-related validation failures.
//
// It inspects the validation tag from the provided validator.FieldError
// and maps it to a corresponding error value.
//
// Supported tags:
//   - "email": the field must be a valid email address
//   - "uuid4": the field must be a valid UUID (version 4)
//
// If the validation tag is not recognized, it returns nil.
func getErrorForStringTag(fe validator.FieldError) error {
	switch fe.Tag() {
	case "email":
		return ErrEmail
	case "uuid4":
		return ErrUUID4
	}

	return nil
}
