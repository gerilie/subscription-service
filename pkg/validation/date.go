package validation

import "github.com/go-playground/validator/v10"

// getErrorMessageForDateTag returns a user-friendly error message for date-related validation errors.
//
// It inspects the validation tag from the provided validator.FieldError
// and maps it to a corresponding message.
//
// Supported tags:
//   - "datetime": the field must follow the "MM-YYYY" format.
//
// If the validation tag is not recognized, getErrorMessageForDateTag returns an empty string.
func getErrorMessageForDateTag(fe validator.FieldError) string {
	switch fe.Tag() {
	case "datetime":
		return "The field must be in the format MM-YYYY"
	}

	return ""
}
