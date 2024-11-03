package controllers

import (
	"DimasPrasetyo/backend-next/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)


func handleError(c *gin.Context, status int, message string) {
	c.JSON(status, gin.H{"error": message})
}

func GetAllAddresses(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var addresses []models.Address
		if err := db.Find(&addresses).Error; err != nil {
			handleError(c, http.StatusInternalServerError, "Could not retrieve addresses")
			return
		}
		c.JSON(http.StatusOK, addresses)
	}
}

func GetAddressByID(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		var address models.Address
		if err := db.First(&address, "id = ?", id).Error; err != nil {
			handleError(c, http.StatusNotFound, "Address not found")
			return
		}
		c.JSON(http.StatusOK, address)
	}
}

func CreateAddress(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var address models.Address
		if err := c.ShouldBindJSON(&address); err != nil {
			handleError(c, http.StatusBadRequest, "Invalid input")
			return
		}
		if err := validate.Struct(address); err != nil {
			handleError(c, http.StatusBadRequest, err.Error())
			return
		}
		if err := db.Create(&address).Error; err != nil {
			handleError(c, http.StatusInternalServerError, "Could not create address")
			return
		}
		c.JSON(http.StatusCreated, address)
	}
}

func UpdateAddress(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		var address models.Address
		if err := db.First(&address, "id = ?", id).Error; err != nil {
			handleError(c, http.StatusNotFound, "Address not found")
			return
		}
		if err := c.ShouldBindJSON(&address); err != nil {
			handleError(c, http.StatusBadRequest, "Invalid input")
			return
		}
		if err := validate.Struct(address); err != nil {
			handleError(c, http.StatusBadRequest, err.Error())
			return
		}
		if err := db.Save(&address).Error; err != nil {
			handleError(c, http.StatusInternalServerError, "Could not update address")
			return
		}
		c.JSON(http.StatusOK, address)
	}
}

func DeleteAddress(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		var address models.Address
		if err := db.First(&address, "id = ?", id).Error; err != nil {
			handleError(c, http.StatusNotFound, "Address not found")
			return
		}
		if err := db.Delete(&address).Error; err != nil {
			handleError(c, http.StatusInternalServerError, "Could not delete address")
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Address deleted successfully"})
	}
}

func GetAddressesByUserID(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := c.Param("userID")
		var addresses []models.Address
		if err := db.Where("user_id = ?", userID).Find(&addresses).Error; err != nil {
			handleError(c, http.StatusInternalServerError, "Could not retrieve addresses")
			return
		}
		c.JSON(http.StatusOK, addresses)
	}
}

func GetAddressesByCity(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		city := c.Param("city")
		var addresses []models.Address
		if err := db.Where("city = ?", city).Find(&addresses).Error; err != nil {
			handleError(c, http.StatusInternalServerError, "Could not retrieve addresses")
			return
		}
		c.JSON(http.StatusOK, addresses)
	}
}

func GetAddressesByCountry(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		country := c.Param("country")
		var addresses []models.Address
		if err := db.Where("country = ?", country).Find(&addresses).Error; err != nil {
			handleError(c, http.StatusInternalServerError, "Could not retrieve addresses")
			return
		}
		c.JSON(http.StatusOK, addresses)
	}
}
