package models

import (
	"errors"
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	ID            string         `gorm:"size:36;not null;uniqueIndex;primary_key"`
	FirstName     string         `gorm:"size:100;not null"`
	Addresses     []Address //has many
	LastName      string         `gorm:"size:100;not null"`
	Email         string         `gorm:"size:100;not null;uniqueIndex"`
	Password      string         `gorm:"size:255;not null"`
	RememberToken string         `gorm:"size:255;not null"`
	Image         string  
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt `gorm:"index"`
}
// HashPassword hashes the user's password
func (u *User) HashPassword(password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}

// CheckPassword checks if the provided password matches the stored hash
func (u *User) CheckPassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
}

// GetUserByEmail retrieves a user by email from the database
func GetUserByEmail(db *gorm.DB, email string, user *User) error {
	result := db.Where("email = ?", email).First(user)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return errors.New("user not found")
	}
	return result.Error
}

// RegisterModels returns a slice of structs that can be used to register models with a gorm.DB.
// The slice returned contains the User, Address, Product, ProductImage and Category models.
// RegisterModels returns a slice of structs that can be used to register models with a gorm.DB.
func RegisterModels() []interface{} {
	return []interface{}{
		&User{},           // Pastikan User dimasukkan di sini
		&Address{},
		&Product{},
		&ProductImage{},
		&Category{},
	}
}
