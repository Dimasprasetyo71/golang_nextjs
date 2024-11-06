package fakers

import (
	"DimasPrasetyo/backend-next/models"
	"DimasPrasetyo/backend-next/utils"
	"math/rand"
	"time"

	"github.com/bxcodec/faker/v4"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func AddressFaker(db *gorm.DB) []models.Address {
	var addresses []models.Address

	var users []models.User
	if err := db.Find(&users).Error; err != nil {
		return addresses
	}

	for i := 0; i < 10; i++ {
		if len(users) == 0 {
			break 
		}

		userID := users[rand.Intn(len(users))].ID 
		cityID := uuid.New().String()
		address := models.Address{
			ID:        uuid.New().String(),
			UserID:    userID,
			Name:      faker.Name(),
			IsPrimary: utils.RandomBool(),
			CityID:    cityID,
			Address1:  utils.RandomStreetAddress(),
			Address2:  utils.RandomSecondaryAddress(),
			Phone:     faker.Phonenumber(),
			Email:     faker.Email(),
			PostCode:  utils.RandomPostCode(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}
		addresses = append(addresses, address)
	}

	return addresses
}
