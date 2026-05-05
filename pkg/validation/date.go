package validation

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

var ErrDatetime = fmt.Errorf("%s in the format 'MM-YYYY'", ValidationPrefix)

// getErrorForDatetimeTag maps a validation error to a user-friendly error.
//
// It inspects the validation tag from the provided validator.FieldError
// and returns a corresponding error value.
//
// Supported tags:
//   - "datetime": the field must match the "MM-YYYY" format.
//
// If the validation tag is not recognized, it returns nil.
func getErrorForDatetimeTag(fe validator.FieldError) error {
	switch fe.Tag() {
	case "datetime":
		return ErrDatetime
	}

	return nil
}
