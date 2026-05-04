package validation

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

func getErrorMessageForNumberTag(fe validator.FieldError) string {
	switch fe.Tag() {
	case "min":
		return fmt.Sprintf(
			"The field must be greater than or equal to %s",
			fe.Param(),
		)
	case "max":
		return fmt.Sprintf(
			"The field must be less than or equal to %s",
			fe.Param(),
		)
	case "gt":
		return fmt.Sprintf(
			"The field must be greater than %s",
			fe.Param(),
		)
	case "gte":
		return fmt.Sprintf(
			"The field must be greater than or equal to %s",
			fe.Param(),
		)
	}

	return ""
}
