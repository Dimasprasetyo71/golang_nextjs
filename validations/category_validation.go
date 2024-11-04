package validations

import (
	"DimasPrasetyo/backend-next/dto"

	"github.com/go-playground/validator/v10"
)

func init() {
	validate = validator.New()
}

func ValidateCategoryRequest(categoryRequest dto.CategoryRequest) error {
	return validate.Struct(categoryRequest)
}

