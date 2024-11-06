package routes

import (
	"DimasPrasetyo/backend-next/controllers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func UserRoutes(router *gin.Engine, db *gorm.DB) {
	// Public routes

	router.GET("/auth/google", controllers.LoginWithGoogle)

	router.GET("/auth/google/callback", controllers.GoogleCallback)
	
	router.GET("/auth/github", controllers.LoginWithGithub)
	
	router.GET("/auth/github/callback", controllers.GithubCallback)
	
	router.GET("/api/users", controllers.GetAllUsers(db))

	router.GET("/api/users/:id", controllers.GetUserByID(db))

	router.GET("/api/users/email/:email", controllers.GetUserByEmail(db))

	router.POST("/api/users/register", controllers.Register(db))

	router.POST("/api/users/login", controllers.Login(db))
	
	router.PUT("/api/users/:id", controllers.UpdateUser(db))

	router.DELETE("/api/users/:id", controllers.DeleteUser(db))

	protected := router.Group("/api/users")

	protected.Use(controllers.ProtectedEndpoint)

}

func AddressRoutes(router *gin.Engine, db *gorm.DB) {

	router.GET("/api/addresses", controllers.GetAllAddresses(db))

	router.GET("/api/addresses/:id", controllers.GetAddressByID(db))

	router.GET("/api/addresses/user/:userID", controllers.GetAddressesByUserID(db))

	router.GET("/api/addresses/city/:city", controllers.GetAddressesByCity(db))

	router.GET("/api/addresses/country/:country", controllers.GetAddressesByCountry(db))

	router.POST("/api/addresses", controllers.CreateAddress(db))

	router.PUT("/api/addresses/:id", controllers.UpdateAddress(db))

	router.DELETE("/api/addresses/:id", controllers.DeleteAddress(db))

}

func CategoryRoutes(router *gin.Engine, db *gorm.DB) {
	router.GET("/api/categories", controllers.GetAllCategories(db))

	router.GET("/api/categories/:id", controllers.GetCategoryByID(db))

	router.POST("/api/categories", controllers.CreateCategory(db))

	router.PUT("/api/categories/:id", controllers.UpdateCategory(db))

	router.DELETE("/api/categories/:id", controllers.DeleteCategory(db))

}

func ProductImageRoutes(router *gin.Engine, db *gorm.DB) {
	router.GET("/api/product-images", controllers.GetAllProductImages(db))

	router.GET("/api/product-images/:id", controllers.GetProductImageByID(db))

	router.POST("/api/product-images", controllers.CreateProductImage(db))

	router.PUT("/api/product-images/:id", controllers.UpdateProductImage(db))

	router.DELETE("/api/product-images/:id", controllers.DeleteProductImage(db))

}

func ProductRoutes(router *gin.Engine, db *gorm.DB) {
	router.GET("/api/products", controllers.GetAllProducts(db))

	router.GET("/api/products/:id", controllers.GetProductByID(db))

	router.POST("/api/products", controllers.CreateProduct(db))

	router.PUT("/api/products/:id", controllers.UpdateProduct(db))

	router.DELETE("/api/products/:id", controllers.DeleteProduct(db))

}