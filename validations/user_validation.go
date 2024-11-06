package validations

import (
	    "github.com/go-playground/validator/v10"

		"DimasPrasetyo/backend-next/dto"
)

var validate = validator.New()


func ValidateUserRequest(userRequest dto.UserRequest) error {
    return validate.Struct(userRequest)
}