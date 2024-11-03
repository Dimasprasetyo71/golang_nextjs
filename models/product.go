package models

import (
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
	"time"
)

type Product struct {
	ID       string `gorm:"size:36;not null;uniqueIndex;primary_key"`
	ParentID string `gorm:"size:36;index"`
	User     User
	UserID   string `gorm:"size:36;index"`
	Name     string `gorm:"size:255;not null"`
	ShortDescription string `gorm:"size:text"`
	Description      string `gorm:"type:text"`
	ProductImages    []ProductImage
	Categories       []Category      `gorm:"many2many:product_categories"`
	Price            decimal.Decimal `gorm:"type:decimal(10,2);not null"`
	Quantity         int             `gorm:"not null"`
	IsAvailable      bool            `gorm:"not null"`
	CreatedAt        time.Time
	UpdatedAt        time.Time
	DeletedAt        gorm.DeletedAt `gorm:"index"`
}
