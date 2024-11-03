package controllers

import (
	"net/http"

	"DimasPrasetyo/backend-next/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetAllProductImages(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var productImages []models.ProductImage
		if err := db.Find(&productImages).Error; err != nil {
			handleError(c, http.StatusInternalServerError, "Could not retrieve product images")
			return
		}
		c.JSON(http.StatusOK, productImages)
	}
}

func GetProductImageByID(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		var productImage models.ProductImage
		if err := db.First(&productImage, "id = ?", id).Error; err != nil {
			handleError(c, http.StatusNotFound, "Product image not found")
			return
		}
		c.JSON(http.StatusOK, productImage)
	}
}

func CreateProductImage(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var productImage models.ProductImage
		if err := c.ShouldBindJSON(&productImage); err != nil {
			handleError(c, http.StatusBadRequest, "Invalid input")
			return
		}
		if err := validate.Struct(productImage); err != nil {
			handleError(c, http.StatusBadRequest, err.Error())
			return
		}
		if err := db.Create(&productImage).Error; err != nil {
			handleError(c, http.StatusInternalServerError, "Could not create product image")
			return
		}
		c.JSON(http.StatusCreated, productImage)
	}
}

func UpdateProductImage(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		var productImage models.ProductImage
		if err := db.First(&productImage, "id = ?", id).Error; err != nil {
			handleError(c, http.StatusNotFound, "Product image not found")
			return
		}
		if err := c.ShouldBindJSON(&productImage); err != nil {
			handleError(c, http.StatusBadRequest, "Invalid input")
			return
		}
		if err := validate.Struct(productImage); err != nil {
			handleError(c, http.StatusBadRequest, err.Error())
			return
		}
		if err := db.Save(&productImage).Error; err != nil {
			handleError(c, http.StatusInternalServerError, "Could not update product image")
			return
		}
		c.JSON(http.StatusOK, productImage)
	}
}

func DeleteProductImage(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		var productImage models.ProductImage
		if err := db.First(&productImage, "id = ?", id).Error; err != nil {
			handleError(c, http.StatusNotFound, "Product image not found")
			return
		}
		if err := db.Delete(&productImage).Error; err != nil {
			handleError(c, http.StatusInternalServerError, "Could not delete product image")
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Product image deleted successfully"})
	}
}

