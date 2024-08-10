package helper

import (
	"fmt"
)

// ValidationError represents an error that occurs during validation
type ValidationError struct {
	Key     string
	Message string
}

// ValidationBuilder struct holds the payload and the accumulated errors
type ValidationBuilder struct {
	Payload map[string]interface{}
	Errors  []ValidationError
}

func NewValidationBuilder(payload map[string]interface{}) *ValidationBuilder {
	return &ValidationBuilder{
		Payload: payload,
		Errors:  []ValidationError{},
	}
}

func (v *ValidationBuilder) ValidateRequiredKeys(requiredKeys []string) *ValidationBuilder {
	for _, key := range requiredKeys {
		if _, ok := v.Payload[key]; !ok {
			v.Errors = append(v.Errors, ValidationError{Key: key, Message: fmt.Sprintf("%s is required", key)})
			return v
		}
	}
	return v
}

func (v *ValidationBuilder) IsEmptyOrNull() *ValidationBuilder {
	for key, value := range v.Payload {
		if value == nil || value == "" {
			v.Errors = append(v.Errors, ValidationError{Key: key, Message: fmt.Sprintf("%s is required", key)})
			return v
		}
	}
	return v
}

func (v *ValidationBuilder) CheckLength(requiredKeys []string) *ValidationBuilder {
	for _, key := range requiredKeys {
		if value, ok := v.Payload[key].(string); ok {
			if len(value) != 16 {
				v.Errors = append(v.Errors, ValidationError{Key: key, Message: fmt.Sprintf("%s must have a length of 16", key)})
				return v
			}
		}
	}
	return v
}

func (v *ValidationBuilder) IsString(requiredKeys []string) *ValidationBuilder {
	for _, key := range requiredKeys {
		if _, ok := v.Payload[key].(string); !ok {
			v.Errors = append(v.Errors, ValidationError{Key: key, Message: fmt.Sprintf("%s must be a string", key)})
			return v
		}
	}
	return v
}

func (v *ValidationBuilder) IsInt(requiredKeys []string) *ValidationBuilder {
	for _, key := range requiredKeys {
		if value, ok := v.Payload[key].(float64); !ok || float64(int(value)) != value {
			v.Errors = append(v.Errors, ValidationError{Key: key, Message: fmt.Sprintf("%s must be an integer", key)})
			return v
		}
	}
	return v
}

// Build returns the accumulated errors after validation
func (v *ValidationBuilder) Build() []ValidationError {
	return v.Errors
}
