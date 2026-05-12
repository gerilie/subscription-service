package validation

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

var (
	// ErrMin indicates that a field value is below the minimum allowed value.
	ErrMin = fmt.Errorf(
		"%w: value is too small",
		ErrValidation,
	)

	// ErrMax indicates that a field value exceeds the maximum allowed value.
	ErrMax = fmt.Errorf(
		"%w: value is too large",
		ErrValidation,
	)

	// ErrGt indicates that a field value must be greater than the specified limit.
	ErrGt = fmt.Errorf(
		"%w: must be greater than the specified value",
		ErrValidation,
	)

	// ErrGte indicates that a field value must be greater than or equal to
	// the specified limit.
	ErrGte = fmt.Errorf(
		"%w: must be greater than or equal to the specified value",
		ErrValidation,
	)
)

// getErrorForNumberTag maps validator numeric tags to user-friendly validation errors.
//
// Supported validator tags:
//   - "min": field value must be greater than or equal to the minimum limit
//   - "max": field value must be less than or equal to the maximum limit
//   - "gt":  field value must be strictly greater than the limit
//   - "gte": field value must be greater than or equal to the limit
//
// Returns nil if the validation tag is unsupported.
func getErrorForNumberTag(fe validator.FieldError) error {
	switch fe.Tag() {
	case "min":
		return MinError(fe.Param())
	case "max":
		return MaxError(fe.Param())
	case "gt":
		return GtError(fe.Param())
	case "gte":
		return GteError(fe.Param())
	}

	return nil
}

// MinError returns ErrMin enriched with the minimum allowed value.
func MinError(value string) error {
	return fmt.Errorf("%w: minimum allowed is %s", ErrMin, value)
}

// MaxError returns ErrMax enriched with the maximum allowed value.
func MaxError(value string) error {
	return fmt.Errorf("%w: maximum allowed is %s", ErrMax, value)
}

// GtError returns ErrGt enriched with the required lower bound.
func GtError(value string) error {
	return fmt.Errorf("%w: must be greater than %s", ErrGt, value)
}

// GteError returns ErrGte enriched with the required inclusive lower bound.
func GteError(value string) error {
	return fmt.Errorf("%w: must be greater than or equal to %s", ErrGte, value)
}
