package validations

import (
	"DimasPrasetyo/backend-next/dto"

	"github.com/go-playground/validator/v10"
)

func init() {
	validate = validator.New()
}

func ValidateAddresRequest(addresRequest dto.AddresRequest) error {
	return validate.Struct(addresRequest)
}