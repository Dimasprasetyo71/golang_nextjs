package test

import (
	"DimasPrasetyo/backend-next/config"
	"DimasPrasetyo/backend-next/models"
	"testing"
)

func TestCreateCategory(t *testing.T) {
	config.ConnectDB()
	db := config.ConnectDB()
	db.Migrator().DropTable(&models.Category{})
	db.AutoMigrate(&models.Category{})

	var category models.Category

	t.Run("Create Category", func(t *testing.T) {
		category = models.Category{
			Name: "Test Category",
			Slug: "test-category",
		}

		if err := db.Create(&category).Error; err != nil {
			t.Errorf("Failed to create category: %v", err)
		}
	})

}

func TestFindCategory(t *testing.T) {
	config.ConnectDB()
	db := config.ConnectDB()
	db.Migrator().DropTable(&models.Category{})
	db.AutoMigrate(&models.Category{})

	var category models.Category

	t.Run("Find Category", func(t *testing.T) {
		category = models.Category{
			Name: "Test Category",
			Slug: "test-category",
		}

		if err := db.Create(&category).Error; err != nil {
			t.Errorf("Failed to create category: %v", err)
		}

		if err := db.First(&category, "id = ?", category.ID).Error; err != nil {
			t.Errorf("Failed to find category: %v", err)
		}
	})
}

func TestUpdateCategory(t *testing.T) {
	config.ConnectDB()
	db := config.ConnectDB()
	db.Migrator().DropTable(&models.Category{})
	db.AutoMigrate(&models.Category{})

	var category models.Category

	t.Run("Update Category", func(t *testing.T) {
		category = models.Category{
			Name: "Test Category",
			Slug: "test-category",
		}

		if err := db.Create(&category).Error; err != nil {
			t.Errorf("Failed to create category: %v", err)
		}

		if err := db.First(&category, "id = ?", category.ID).Error; err != nil {
			t.Errorf("Failed to find category: %v", err)

		}
	})

}

func TestDeleteCategory(t *testing.T) {
	config.ConnectDB()
	db := config.ConnectDB()
	db.Migrator().DropTable(&models.Category{})
	db.AutoMigrate(&models.Category{})

	var category models.Category

	t.Run("Delete Category", func(t *testing.T) {
		category = models.Category{
			Name: "Test Category",
			Slug: "test-category",
		}

		if err := db.Create(&category).Error; err != nil {
			t.Errorf("Failed to create category: %v", err)
		}

		if err := db.First(&category, "id = ?", category.ID).Error; err != nil {
			t.Errorf("Failed to find category: %v", err)

		}

		if err := db.Delete(&category).Error; err != nil {
			t.Errorf("Failed to delete category: %v", err)

		}

		if err := db.First(&category, "id = ?", category.ID).Error; err == nil {
			t.Errorf("Failed to delete category: %v", err)
		}
	})
}
