package validation

import (
	"errors"
	"fmt"
	validation "github.com/go-ozzo/ozzo-validation"
	"strings"
)

func ValidationHandler[T any](data *T, rules []*validation.FieldRules) error {
	err := validation.ValidateStruct(data, rules...)
	if err != nil {
		return err
	}
	return nil
}

func HandleValidationError(err error) string {
	var errorMessages []string

	var validationErrors validation.Errors
	if errors.As(err, &validationErrors) {
		for field, rule := range validationErrors {
			errorMessages = append(errorMessages, fmt.Sprintf("%s: %s", field, rule.Error()))
		}
	}

	return strings.Join(errorMessages, "\n")
}
