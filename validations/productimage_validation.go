package validations

import (
	"DimasPrasetyo/backend-next/dto"

	"github.com/go-playground/validator/v10"
)
func init() {
	validate = validator.New()
}

func ValidateProductImageRequest(productImageRequest dto.ProductImageRequest) error {
	return validate.Struct(productImageRequest)
}

func ValidateProductImageUpdateRequest(productImageUpdateRequest dto.ProductImageUpdateRequest) error {
	return validate.Struct(productImageUpdateRequest)
}
