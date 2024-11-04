package repositories

import (
	"DimasPrasetyo/backend-next/models"
	"errors"

	"gorm.io/gorm"
)

type CategoryRepository struct {
	db *gorm.DB
}
type CategoryRepositoryInterface interface {
	FindAll() ([]models.Category, error)

	FindByID(id string) (*models.Category, error)

	CreateCategory(category *models.Category) error

	UpdateCategory(category *models.Category) error

	DeleteCategory(id string) error

	FindByName(name string) (*models.Category, error)

	FindBySlug(slug string) (*models.Category, error)

	FindByParentID(parentID string) ([]models.Category, error)

	FindAllByParentID(parentID string) ([]models.Category, error)
}

func NewCategoryRepository(db *gorm.DB) CategoryRepository {
	return CategoryRepository{db: db}
}

func (r CategoryRepository) FindAll() ([]models.Category, error) {

	var categories []models.Category

	if err := r.db.Find(&categories).Error; err != nil {

		return nil, err

	}

	return categories, nil

}

func (r CategoryRepository) FindByID(id string) (*models.Category, error) {

	var category models.Category

	if err := r.db.First(&category, "id = ?", id).Error; err != nil {

		if errors.Is(err, gorm.ErrRecordNotFound) {

			return nil, errors.New("category not found")

		}

		return nil, err

	}

	return &category, nil
}

func (r CategoryRepository) CreateCategory(category *models.Category) error {

	if err := r.db.Create(&category).Error; err != nil {

		return err

	}

	return nil

}

func (r CategoryRepository) UpdateCategory(category *models.Category) error {

	if err := r.db.Save(&category).Error; err != nil {

		return err

	}

	return nil

}

func (r CategoryRepository) DeleteCategory(id string) error {

	if err := r.db.Delete(&models.Category{}, "id = ?", id).Error; err != nil {

		return err

	}

	return nil

}

func (r CategoryRepository) FindByName(name string) (*models.Category, error) {

	var category models.Category

	if err := r.db.First(&category, "name = ?", name).Error; err != nil {

		if errors.Is(err, gorm.ErrRecordNotFound) {

			return nil, errors.New("category not found")

		}

		return nil, err

	}

	return &category, nil

}

func (r CategoryRepository) FindBySlug(slug string) (*models.Category, error) {

	var category models.Category

	if err := r.db.First(&category, "slug = ?", slug).Error; err != nil {

		if errors.Is(err, gorm.ErrRecordNotFound) {

			return nil, errors.New("category not found")

		}

		return nil, err

	}

	return &category, nil

}

func (r CategoryRepository) FindByParentID(parentID string) ([]models.Category, error) {

	var categories []models.Category

	if err := r.db.Find(&categories, "parent_id = ?", parentID).Error; err != nil {

		return nil, err

	}

	return categories, nil

}

func (r CategoryRepository) FindAllByParentID(parentID string) ([]models.Category, error) {

	var categories []models.Category

	if err := r.db.Find(&categories, "parent_id = ?", parentID).Error; err != nil {

		return nil, err

	}

	return categories, nil

}

func (r CategoryRepository) FindByParentIDAndSlug(parentID string, slug string) (*models.Category, error) {

	var category models.Category

	if err := r.db.First(&category, "parent_id = ? AND slug = ?", parentID, slug).Error; err != nil {

		if errors.Is(err, gorm.ErrRecordNotFound) {

			return nil, errors.New("category not found")

		}

		return nil, err

	}

	return &category, nil

}