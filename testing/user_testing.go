package test

import (
	"DimasPrasetyo/backend-next/config"
	"DimasPrasetyo/backend-next/models"
	"testing"
)

func Userall(t *testing.T) {
	config.ConnectDB()
	db := config.ConnectDB()
	db.Migrator().DropTable(&models.User{})
	db.AutoMigrate(&models.User{})

	var user models.User

	t.Run("Create User", func(t *testing.T) {
		user = models.User{
			FirstName: "Dimas",
			LastName:  "Prasetyo",
			Email:     "nLbP0@example.com",
			Password:  "password",
		}

		if err := db.Create(&user).Error; err != nil {
			t.Errorf("Failed to create user: %v", err)
		}
	})
}


