package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"DimasPrasetyo/backend-next/models"
	"gorm.io/gorm"
)

func GetAllProducts(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var products []models.Product
		if err := db.Find(&products).Error; err != nil {
			handleError(c, http.StatusInternalServerError, "Could not retrieve products")
			return
		}
		c.JSON(http.StatusOK, products)
	}
}

func GetProductByID(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		var product models.Product
		if err := db.First(&product, "id = ?", id).Error; err != nil {
			handleError(c, http.StatusNotFound, "Product not found")
			return
		}
		c.JSON(http.StatusOK, product)
	}
}


func CreateProduct(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var product models.Product
		if err := c.ShouldBindJSON(&product); err != nil {
			handleError(c, http.StatusBadRequest, "Invalid request payload")
			return
		}
		if err := validate.Struct(product); err != nil {
			handleError(c, http.StatusBadRequest, err.Error())
			return
		}
		if err := db.Create(&product).Error; err != nil {
			handleError(c, http.StatusInternalServerError, "Failed to create product")
			return
		}
		c.JSON(http.StatusCreated, product)
	}
}

func UpdateProduct(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		var product models.Product
		if err := db.First(&product, "id = ?", id).Error; err != nil {
			handleError(c, http.StatusNotFound, "Product not found")
			return
		}
		if err := c.ShouldBindJSON(&product); err != nil {
			handleError(c, http.StatusBadRequest, "Invalid request payload")
			return
		}
		if err := validate.Struct(product); err != nil {
			handleError(c, http.StatusBadRequest, err.Error())
			return
		}
		if err := db.Save(&product).Error; err != nil {
			handleError(c, http.StatusInternalServerError, "Failed to update product")
			return
		}
		c.JSON(http.StatusOK, product)
	}
}

func DeleteProduct(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		var product models.Product
		if err := db.First(&product, "id = ?", id).Error; err != nil {
			handleError(c, http.StatusNotFound, "Product not found")
			return
		}
		if err := db.Delete(&product).Error; err != nil {
			handleError(c, http.StatusInternalServerError, "Failed to delete product")
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Product deleted successfully"})
	}
}
	