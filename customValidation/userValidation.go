package customValidation

import (
	"api/go/helper"
)

func ValidatePayload(payload map[string]interface{}) []helper.ValidationError {
	validationErrors := helper.NewValidationBuilder(payload).
		ValidateRequiredKeys([]string{"id", "firstName", "middleName", "lastName"}).
		IsEmptyOrNull().
		IsString([]string{"firstName", "middleName", "lastName"}).
		Build()

		// Convert []helper.ValidationError to []string

	return validationErrors
}
