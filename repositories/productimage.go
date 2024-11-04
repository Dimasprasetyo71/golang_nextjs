package repositories

import (
	"DimasPrasetyo/backend-next/models"
	"errors"
	"gorm.io/gorm"
)

type ProductImageRepository interface {

	FindAll() ([]models.ProductImage, error)

	FindByID(id string) (*models.ProductImage, error)

	CreateProductImage(productImage *models.ProductImage) error

	UpdateProductImage(productImage *models.ProductImage) error

	DeleteProductImage(productImage *models.ProductImage) error

}

type productImageRepository struct {
	db *gorm.DB
}

func NewProductImageRepository(db *gorm.DB) ProductImageRepository {
	return &productImageRepository{db}
}

func (r *productImageRepository) FindAll() ([]models.ProductImage, error) {

	var productImages []models.ProductImage

	if err := r.db.Find(&productImages).Error; err != nil {

		return nil, err

	}

	return productImages, nil

}

func (r *productImageRepository) FindByID(id string) (*models.ProductImage, error) {

	var productImage models.ProductImage

	if err := r.db.First(&productImage, "id = ?", id).Error; err != nil {

		if errors.Is(err, gorm.ErrRecordNotFound) {

			return nil, errors.New("product image not found")

		}

		return nil, err

	}

	return &productImage, nil

}

func (r *productImageRepository) CreateProductImage(productImage *models.ProductImage) error {

	if err := r.db.Create(productImage).Error; err != nil {

		return err

	}

	return nil

}

func (r *productImageRepository) UpdateProductImage(productImage *models.ProductImage) error {

	if err := r.db.Save(productImage).Error; err != nil {

		return err

	}

	return nil

}

func (r *productImageRepository) DeleteProductImage(productImage *models.ProductImage) error {

	if err := r.db.Delete(productImage).Error; err != nil {

		return err

	}

	return nil

}
