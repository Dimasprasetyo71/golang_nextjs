package repositories

import (
    "errors"
    "gorm.io/gorm"
	
	"DimasPrasetyo/backend-next/models"
)

type addresRepository struct {
	db *gorm.DB
}

func NewAddresRepository(db *gorm.DB) *addresRepository {
	return &addresRepository{db: db}
}

func (r *addresRepository) FindAll() ([]models.Address, error) {
	var addres []models.Address
	if err := r.db.Find(&addres).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("addres not found")
		}
		return nil, err
	}
	return addres, nil
}

func (r *addresRepository) FindByID(id string) (*models.Address, error) {
	var addres models.Address
	if err := r.db.First(&addres, "id = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("addres not found")
		}
		return nil, err
	}
	return &addres, nil
}

func (r *addresRepository) CreateAddres(addres *models.Address) error {
	return r.db.Create(addres).Error
}

func (r *addresRepository) UpdateAddres(addres *models.Address) error {

	return r.db.Save(addres).Error
}

func (r *addresRepository) DeleteAddres(addres *models.Address) error {
	return r.db.Delete(addres).Error
}

func (r *addresRepository) FindByUserID(id string) ([]models.Address, error) {
	var addres []models.Address
	if err := r.db.Where("user_id = ?", id).Find(&addres).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("addres not found")
		}
		return nil, err
	}
	return addres, nil
}

func (r *addresRepository) FindByUserIDAndIsPrimary(id string) (*models.Address, error) {
	var addres models.Address
	if err := r.db.Where("user_id = ? AND is_primary = ?", id, true).First(&addres).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("addres not found")
		}
		return nil, err
	}
	return &addres, nil
}

func (r *addresRepository) FindByUserIDAndIsNotPrimary(id string) ([]models.Address, error) {
	var addres []models.Address
	if err := r.db.Where("user_id = ? AND is_primary = ?", id, false).Find(&addres).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("addres not found")
		}
		return nil, err
	}
	return addres, nil
}

func (r *addresRepository) FindByUserIDAndIsPrimaryAndIsNotPrimary(id string) (*models.Address, error) {
	var addres models.Address
	if err := r.db.Where("user_id = ? AND is_primary = ? AND is_primary = ?", id, true, false).First(&addres).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("addres not found")
		}
		return nil, err
	}
	return &addres, nil
}