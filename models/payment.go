package models

import (
	"time"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type Payment struct {
	ID           string          `gorm:"size:36;not null;uniqueIndex;primary_key"`
	OrderID      string          `gorm:"size:36;index"`       // For tracking purposes with Midtrans
	UserID       string          `gorm:"size:36;index"`       // User associated with the payment
	Amount       decimal.Decimal `gorm:"type:decimal(10,2);not null"`
	Status       string          `gorm:"size:50;not null"`    
	PaymentType  string          `gorm:"size:50;not null"`    
	TransactionID string         `gorm:"size:50;uniqueIndex"` // Unique transaction ID from Midtrans
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt `gorm:"index"`
}
