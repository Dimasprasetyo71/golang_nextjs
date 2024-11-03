package repositories

import (

	"gorm.io/gorm"

	"DimasPrasetyo/backend-next/models"
)



type ProductRepository interface {
	FindAll() ([]models.Product, error)
	FindByID(id string) (*models.Product, error)
	FindByCategoryID(id string) ([]models.Product, error)
	FindByUserID(id string) ([]models.Product, error)
	CreateProduct(product *models.Product) error
	UpdateProduct(product *models.Product) error
	DeleteProduct(product *models.Product) error
}

type productRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &productRepository{db}
}

func (repo *productRepository) FindAll() ([]models.Product, error) {
	var products []models.Product
	err := repo.db.Find(&products).Error
	return products, err
}

func (repo *productRepository) FindByID(id string) (*models.Product, error) {
	var product models.Product
	err := repo.db.First(&product, "id = ?", id).Error
	return &product, err
}

func (repo *productRepository) FindByCategoryID(id string) ([]models.Product, error) {
	var products []models.Product
	err := repo.db.Where("category_id = ?", id).Find(&products).Error
	return products, err
}

func (repo *productRepository) FindByUserID(id string) ([]models.Product, error) {
	var products []models.Product
	err := repo.db.Where("user_id = ?", id).Find(&products).Error
	return products, err
}

func (repo *productRepository) CreateProduct(product *models.Product) error {
	err := repo.db.Create(&product).Error
	return err
}

func (repo *productRepository) UpdateProduct(product *models.Product) error {
	err := repo.db.Save(&product).Error
	return err
}

func (repo *productRepository) DeleteProduct(product *models.Product) error {
	err := repo.db.Delete(&product).Error
	return err
}

