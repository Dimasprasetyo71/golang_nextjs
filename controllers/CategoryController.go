package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"DimasPrasetyo/backend-next/models"

	"gorm.io/gorm"
	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

func CreateCategory(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var category models.Category
		if err := c.ShouldBindJSON(&category); err != nil {
			handleError(c, http.StatusBadRequest, "Invalid input")
			return
		}

		if err := validate.Struct(category); err != nil {
			handleError(c, http.StatusBadRequest, err.Error())
			return
		}

		if err := db.Create(&category).Error; err != nil {
			handleError(c, http.StatusInternalServerError, "Could not create category")
			return
		}

		c.JSON(http.StatusCreated, category)
	}
}

func GetAllCategories(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var categories []models.Category
		if err := db.Find(&categories).Error; err != nil {
			handleError(c, http.StatusInternalServerError, "Could not retrieve categories")
			return
		}

		c.JSON(http.StatusOK, categories)
	}
}

func GetCategoryByID(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		var category models.Category
		if err := db.First(&category, "id = ?", id).Error; err != nil {
			handleError(c, http.StatusNotFound, "Category not found")
			return
		}

		c.JSON(http.StatusOK, category)
	}
}

func UpdateCategory(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		var category models.Category
		if err := db.First(&category, "id = ?", id).Error; err != nil {
			handleError(c, http.StatusNotFound, "Category not found")
			return
		}

		if err := c.ShouldBindJSON(&category); err != nil {
			handleError(c, http.StatusBadRequest, "Invalid input")
			return
		}

		if err := validate.Struct(category); err != nil {
			handleError(c, http.StatusBadRequest, err.Error())
			return
		}

		if err := db.Save(&category).Error; err != nil {
			handleError(c, http.StatusInternalServerError, "Could not update category")
			return
		}

		c.JSON(http.StatusOK, category)
	}
}

func DeleteCategory(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		var category models.Category
		if err := db.First(&category, "id = ?", id).Error; err != nil {
			handleError(c, http.StatusNotFound, "Category not found")
			return
		}

		if err := db.Delete(&category).Error; err != nil {
			handleError(c, http.StatusInternalServerError, "Could not delete category")
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Category deleted successfully"})
	}
}

