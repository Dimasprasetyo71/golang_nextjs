package test

import (
	"DimasPrasetyo/backend-next/config"
	"DimasPrasetyo/backend-next/models"
	"testing"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

var db *gorm.DB

func setup() {
	db = config.ConnectDB()
	if db == nil {
		panic("Failed to connect to the database")
	}
	db.Migrator().DropTable(&models.Address{})
	db.AutoMigrate(&models.Address{})
}

func tearDown() {
	db.Migrator().DropTable(&models.Address{})
}

func createTestAddress() models.Address {
	return models.Address{
		UserID:    uuid.NewString(),
		Email:     "nLbP0@example.com",
		Phone:     "08123456789",
		PostCode:  "12345",
		Address1:  "Jl. Jendral Sudirman",
		Address2:  "No. 1",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func TestCreateAddress(t *testing.T) {
	setup()
	defer tearDown()

	address := createTestAddress()
	if err := db.Create(&address).Error; err != nil {
		t.Errorf("Failed to create address: %v", err)
	}
}

func TestUpdateAddress(t *testing.T) {
	setup()
	defer tearDown()

	address := createTestAddress()
	if err := db.Create(&address).Error; err != nil {
		t.Errorf("Failed to create address: %v", err)
	}

	address.Address1 = "Updated Address"
	if err := db.Save(&address).Error; err != nil {
		t.Errorf("Failed to update address: %v", err)
	}

	var updatedAddress models.Address
	if err := db.First(&updatedAddress, "id = ?", address.ID).Error; err != nil {
		t.Errorf("Failed to find updated address: %v", err)
	}
}

func TestDeleteAddress(t *testing.T) {
	setup()
	defer tearDown()

	address := createTestAddress()
	if err := db.Create(&address).Error; err != nil {
		t.Errorf("Failed to create address: %v", err)
	}

	if err := db.Delete(&address).Error; err != nil {
		t.Errorf("Failed to delete address: %v", err)
	}

	if err := db.First(&address, "id = ?", address.ID).Error; err == nil {
		t.Errorf("Expected address to be deleted, but found: %v", address)
	}
}

func TestFindAddress(t *testing.T) {
	setup()
	defer tearDown()

	address := createTestAddress()
	if err := db.Create(&address).Error; err != nil {
		t.Errorf("Failed to create address: %v", err)
	}

	var foundAddress models.Address
	if err := db.First(&foundAddress, "id = ?", address.ID).Error; err != nil {
		t.Errorf("Failed to find address: %v", err)
	}
}
