package validations

import (
	"DimasPrasetyo/backend-next/dto"

	"github.com/go-playground/validator/v10"
)

func init() {
	validate = validator.New()
}

func ValidateProductRequest(productRequest dto.ProductRequest) error {
	return validate.Struct(productRequest)
}


