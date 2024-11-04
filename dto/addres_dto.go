package dto

type AddresRequest struct {
	Name      string `json:"name" validate:"required,min=2,max=100"`
	Address1  string `json:"address1" validate:"required,min=2,max=100"`
	Address2  string `json:"address2" validate:"required,min=2,max=100"`
	Email     string `json:"email" validate:"required,email"`
	Phone     string `json:"phone" validate:"required"`
	CityID    string `json:"city_id" validate:"required"`
	PostCode  string `json:"post_code" validate:"required"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	IsPrimary bool `json:"is_primary"`


}

type AddresResponse struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Address1  string `json:"address1"`
	Address2  string `json:"address2"`
	Phone     string `json:"phone"`
	CityID    string `json:"city_id"`
	PostCode  string `json:"post_code"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	IsPrimary bool `json:"is_primary"`
	Email     string `json:"email"`
}