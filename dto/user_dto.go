// dto/user_dto.go
package dto

type UserRequest struct {
    FirstName string `json:"first_name" validate:"required,min=2,max=100"`
    LastName  string `json:"last_name" validate:"required,min=2,max=100"`
    Email     string `json:"email" validate:"required,email"`
    Password  string `json:"password" validate:"required,min=6"`
}

type UserResponse struct {
    ID        string `json:"id"`
    FirstName string `json:"first_name"`
    LastName  string `json:"last_name"`
    Email     string `json:"email"`
    Image     string `json:"image"`
}
