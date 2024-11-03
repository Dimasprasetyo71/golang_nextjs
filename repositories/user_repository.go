// repositories/user_repository.go
package repositories

import (
    "errors"
    "gorm.io/gorm"
	
	"DimasPrasetyo/backend-next/models"
)

type UserRepository interface {
    FindByEmail(email string) (*models.User, error)
    CreateUser(user *models.User) error
    UpdateUser(user *models.User) error
}

type userRepository struct {
    db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
    return &userRepository{db}
}

func (r *userRepository) FindByEmail(email string) (*models.User, error) {
    var user models.User
    if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
        if errors.Is(err, gorm.ErrRecordNotFound) {
            return nil, errors.New("user not found")
        }
        return nil, err
    }
    return &user, nil
}

func (r *userRepository) CreateUser(user *models.User) error {
    return r.db.Create(user).Error
}

func (r *userRepository) UpdateUser(user *models.User) error {
    return r.db.Save(user).Error
}
