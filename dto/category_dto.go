package dto


type CategoryRequest struct {
	Name string `json:"name" validate:"required"`
	Slug string
}

type CategoryResponse struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Slug      string `json:"slug"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
