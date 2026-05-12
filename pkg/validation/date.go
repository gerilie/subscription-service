package validation

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

// ErrDatetime indicates that a field value
// does not match the required date format.
var ErrDatetime = fmt.Errorf(
	"%w: must be in the format MM-YYYY",
	ErrValidation,
)

// ErrDateOrder indicates that a start date is after the end date.
var ErrDateOrder = fmt.Errorf(
	"%w: start date must be before end date",
	ErrValidation,
)

// getErrorForDatetimeTag maps validator datetime-related tags
// to user-friendly validation errors.
//
// Supported validator tags:
//   - "datetime": field value must match the "MM-YYYY" format
//
// Returns nil if the validation tag is unsupported.
func getErrorForDatetimeTag(fe validator.FieldError) error {
	if fe.Tag() == "datetime" {
		return ErrDatetime
	}

	return nil
}
