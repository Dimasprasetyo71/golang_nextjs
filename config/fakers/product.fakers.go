package fakers

import (
	"DimasPrasetyo/backend-next/models"
	"DimasPrasetyo/backend-next/utils"
	"github.com/bxcodec/faker/v4"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

func ProductFaker(db *gorm.DB) ([]models.Product, []string) {
	var products []models.Product
	var userIDs []string

	// Retrieve existing user IDs
	db.Model(&models.User{}).Pluck("id", &userIDs)

	for i := 0; i < 10; i++ {
		product := models.Product{
			ID:               uuid.New().String(),
			UserID:           userIDs[i%len(userIDs)], // Use existing user IDs
			Name:             faker.Name(),
			ShortDescription: faker.Name(),
			Description:      faker.Name(),
			Price:            utils.RandomDecimal(100, 1000),
			Quantity:         utils.RandomInt(1, 10),
			IsAvailable:      utils.RandomBool(),
			CreatedAt:        time.Now(),
			UpdatedAt:        time.Now(),
		}
		products = append(products, product)
	}
	return products, getProductIDs(products) // Return product IDs
}

func getProductIDs(products []models.Product) []string {
	var ids []string
	for _, product := range products {
		ids = append(ids, product.ID)
	}
	return ids
}

func ProductImageFaker(db *gorm.DB, productIDs []string) []models.ProductImage {
	var productImages []models.ProductImage

	for i := 0; i < 10; i++ {
		productImage := models.ProductImage{
			ID:         uuid.New().String(),
			ProductID:  productIDs[i%len(productIDs)], // Use existing product IDs
			Path:       faker.Name(),
			ExtraLarge: faker.Name(),
			Large:      faker.Name(),
			Medium:     faker.Name(),
			Small:      faker.Name(),
			CreatedAt:  time.Now(),
			UpdatedAt:  time.Now(),
		}
		productImages = append(productImages, productImage)
	}
	return productImages
}

func CategoryFaker(db *gorm.DB) []models.Category {
	var categories []models.Category
	for i := 0; i < 10; i++ {
		category := models.Category{
			ID:        uuid.New().String(),
			ParentID:  uuid.New().String(),
			Name:      faker.Name(),
			Slug:      faker.Name(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}
		categories = append(categories, category)
	}
	return categories
}
