package validation

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

// getErrorForNumberTag returns a user-friendly error for numeric validation failures.
//
// It inspects the validation tag from the provided validator.FieldError
// and maps it to a corresponding error value.
//
// Supported tags:
//   - "min": the field must be greater than or equal to the specified value
//   - "max": the field must be less than or equal to the specified value
//   - "gt":  the field must be greater than the specified value
//   - "gte": the field must be greater than or equal to the specified value
//
// If the validation tag is not recognized, it returns nil.
func getErrorForNumberTag(fe validator.FieldError) error {
	switch fe.Tag() {
	case "min":
		return ErrMin(fe.Param())
	case "max":
		return ErrMax(fe.Param())
	case "gt":
		return ErrGt(fe.Param())
	case "gte":
		return ErrGte(fe.Param())
	}

	return nil
}

var ErrMin = ErrGte

func ErrMax(value string) error {
	return fmt.Errorf(
		"%s less than or equal to %s",
		ValidationPrefix,
		value,
	)
}

func ErrGt(value string) error {
	return fmt.Errorf(
		"%s greater than %s",
		ValidationPrefix,
		value,
	)
}

func ErrGte(value string) error {
	return fmt.Errorf(
		"%s greater than or equal to %s",
		ValidationPrefix,
		value,
	)
}
