package services

import (
	"DimasPrasetyo/backend-next/dto"
	"DimasPrasetyo/backend-next/models"
	"DimasPrasetyo/backend-next/repositories"
)

type CategoryService interface {
	CreateCategory(category dto.CategoryRequest) (models.Category, error)
}

type CategoryServices struct {
	categoryRepository repositories.CategoryRepository
}

func NewCategoryService(categoryRepository repositories.CategoryRepository) CategoryService {
	return &CategoryServices{
		categoryRepository: categoryRepository,
	}
}

func (s *CategoryServices) CreateCategory(category dto.CategoryRequest) (models.Category, error) {

	categoryModel := models.Category{
		Name:        category.Name,
		Slug:        category.Slug,
	}

	if err := s.categoryRepository.CreateCategory(&categoryModel); err != nil {

		return models.Category{}, err

	}

	return categoryModel, nil

}

func (s *CategoryServices) UpdateCategory(category dto.CategoryRequest) (models.Category, error) {

	categoryModel := models.Category{
		Name:        category.Name,
		Slug:        category.Slug,
	}

	if err := s.categoryRepository.UpdateCategory(&categoryModel); err != nil {

		return models.Category{}, err

	}

	return categoryModel, nil

}

func (s *CategoryServices) DeleteCategory(id string) error {

	if err := s.categoryRepository.DeleteCategory(id); err != nil {

		return err

	}

	return nil

}

func (s *CategoryServices) FindBySlug(slug string) (models.Category, error) {

	category, err := s.categoryRepository.FindBySlug(slug)

	if err != nil {

		return models.Category{}, err

	}

	return *category, nil

}

func (s *CategoryServices) FindByName(name string) (models.Category, error) {

	category, err := s.categoryRepository.FindByName(name)

	if err != nil {

		return models.Category{}, err

	}

	return *category, nil

}

func (s *CategoryServices) FindByID(id string) (models.Category, error) {

	category, err := s.categoryRepository.FindByID(id)

	if err != nil {

		return models.Category{}, err

	}

	return *category, nil

}

func (s *CategoryServices) GetAllCategory() ([]models.Category, error) {

	categories, err := s.categoryRepository.FindAll()

	if err != nil {

		return nil, err

	}

	return categories, nil

}

func (s *CategoryServices) FindByParentID(parentID string) ([]models.Category, error) {

	categories, err := s.categoryRepository.FindByParentID(parentID)

	if err != nil {

		return nil, err

	}

	return categories, nil

}

func (s *CategoryServices) FindAllByParentID(parentID string) ([]models.Category, error) {

	categories, err := s.categoryRepository.FindAllByParentID(parentID)

	if err != nil {

		return nil, err

	}

	return categories, nil

}

func (s *CategoryServices) FindByParentIDAndSlug(parentID string, slug string) (*models.Category, error) {

	category, err := s.categoryRepository.FindByParentIDAndSlug(parentID, slug)

	if err != nil {

		return nil, err

	}

	return category, nil

}

func (s *CategoryServices) FindAll() ([]models.Category, error) {

	categories, err := s.categoryRepository.FindAll()

	if err != nil {

		return nil, err

	}

	return categories, nil

}

