package dto

type MidtransResponse struct {
	ID        string `json:"id"`
	OrderID   string `json:"order_id"`
	UserID    string `json:"user_id"`
	Amount    int `json:"amount"`
	Status    string `json:"status"`
	PaymentType string `json:"payment_type"`
	TransactionID string `json:"transaction_id"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	DeletedAt string `json:"deleted_at"`
}