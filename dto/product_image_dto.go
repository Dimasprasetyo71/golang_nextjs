package dto

type ProductImageRequest struct {
	ID        string
	ProductID string `json:"product_id" validate:"required"`
	Image     string `json:"image" validate:"required"`
	Path      string `json:"path" validate:"required"`

	ExtraLarge string `json:"extra_large"`
	Large      string `json:"large"`
	Medium     string `json:"medium"`
	Small      string `json:"small"`

	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type ProductImageResponse struct {
	ID        string `json:"id"`
	ProductID string `json:"product_id"`
	Image     string `json:"image"`
	Path      string `json:"path"`

	ExtraLarge string `json:"extra_large"`
	Large      string `json:"large"`
	Medium     string `json:"medium"`
	Small      string `json:"small"`

	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type ProductImageUpdateRequest struct {
	ID        string
	ProductID string `json:"product_id" validate:"required"`
	Image     string `json:"image" validate:"required"`
	Path      string `json:"path" validate:"required"`
}

type ProductImageUpdateResponse struct {
	ID        string `json:"id"`
	ProductID string `json:"product_id"`
	Image     string `json:"image"`
	Path      string `json:"path"`
}