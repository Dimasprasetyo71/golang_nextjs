package dto


type ProductRequest struct {
	Name        string  `json:"name" validate:"required"`
	Description string  `json:"description" validate:"required"`
	Price       float64 `json:"price" validate:"required"`
	CategoryID  int     `json:"category_id" validate:"required"`
}

type ProductResponse struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       float64 `json:"price"`
	CategoryID  int     `json:"category_id"`
	ShortDescription string `json:"short_description"`
	Quantity         int    `json:"quantity"`
	IsAvailable      bool   `json:"is_available"`
}