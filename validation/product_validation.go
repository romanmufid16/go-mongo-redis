package validation

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/romanmufid16/go-mongo-redis/model"
)

func ProductValidation(product *model.Product) []*validation.FieldRules {
	return []*validation.FieldRules{
		// Validasi untuk field Name
		validation.Field(&product.Name,
			validation.Required.Error("Name is required"),
			validation.Length(1, 100).Error("Name must be between 1 and 100 characters"),
		),
		// Validasi untuk field Price
		validation.Field(&product.Price,
			validation.Required.Error("Price is required"),
			validation.Min(int64(1)).Error("Price must be greater than 0"),
		),
		// Validasi untuk field Category
		validation.Field(&product.Category,
			validation.Required.Error("Category is required"),
		),
	}
}
