package validation

import (
	validation "github.com/go-ozzo/ozzo-validation"
)

func ValidationHandler[T any](data *T, rules []*validation.FieldRules) error {
	err := validation.ValidateStruct(data, rules...)
	if err != nil {
		return err
	}
	return nil
}

//func Validate[T any](schema []validation.Rule, data T) error {
//	err := validation.Validate(data, schema...)
//	if err != nil {
//		return err
//	}
//	return nil
//}
