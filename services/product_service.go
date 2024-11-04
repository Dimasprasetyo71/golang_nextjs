package services

import (
	"DimasPrasetyo/backend-next/dto"
	"DimasPrasetyo/backend-next/models"
	"DimasPrasetyo/backend-next/repositories"
)

type ProductService interface {
	GetAllProduct() ([]models.Product, error)
	GetProductById(id string) (models.Product, error)
	CreateProduct(productRequest dto.ProductRequest) (models.Product, error)
	UpdateProduct(id string, productRequest dto.ProductRequest) (models.Product, error)
	DeleteProduct(id string) error
}

type productService struct {
	productRepository repositories.ProductRepository
}

// CreateProduct implements ProductService.
func (p *productService) CreateProduct(productRequest dto.ProductRequest) (models.Product, error) {
	panic("unimplemented")
}

// DeleteProduct implements ProductService.
func (p *productService) DeleteProduct(id string) error {
	panic("unimplemented")
}

// GetAllProduct implements ProductService.
func (p *productService) GetAllProduct() ([]models.Product, error) {
	panic("unimplemented")
}

// GetProductById implements ProductService.
func (p *productService) GetProductById(id string) (models.Product, error) {
	panic("unimplemented")
}

// UpdateProduct implements ProductService.
func (p *productService) UpdateProduct(id string, productRequest dto.ProductRequest) (models.Product, error) {
	panic("unimplemented")
}

func NewProductService(productRepository repositories.ProductRepository) ProductService {
	return &productService{productRepository}
}
