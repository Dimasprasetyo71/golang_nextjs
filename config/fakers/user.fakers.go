package fakers

import (
	"DimasPrasetyo/backend-next/models"
	"DimasPrasetyo/backend-next/utils"
	"time"

	"github.com/bxcodec/faker/v4"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"golang.org/x/crypto/bcrypt"
)

func UserFaker(db *gorm.DB) []models.User {
	var users []models.User
	for i := 0; i < 10; i++ {
		password := faker.Password()
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if err != nil {
			panic(err) 
		}

		user := models.User{
			ID:            uuid.New().String(),
			FirstName:     faker.Name(),
			LastName:      faker.Name(),
			Email:         faker.Email(),
			Password:      string(hashedPassword), 
			CreatedAt:     time.Now(),
			UpdatedAt:     time.Now(),
		}

		token, err := utils.GenerateToken(user.ID)
		if err != nil {
			panic(err) 
		}
		user.RememberToken = token 

		if err := db.Create(&user).Error; err != nil {
			panic(err)
		}
		users = append(users, user)
	}
	return users
}
