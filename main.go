package main

import (
	"DimasPrasetyo/backend-next/config"
	"DimasPrasetyo/backend-next/middleware"
	"DimasPrasetyo/backend-next/routes"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	port := os.Getenv("GOLANG_PORT")
	if port == "" {
		log.Fatal("$GOLANG_PORT must be set")
	}

	// Initialize Gin router
	router := gin.New()
	router.Use(cors.Default())
	router.Use(middleware.ErrorHandlerMiddleware())
	router.Use(middleware.RateLimitMiddleware())
	// Connect to the database
	db := config.ConnectDB()
	gin.SetMode(gin.ReleaseMode)


	// Set up routes
	routes.UserRoutes(router, db)
	routes.AddressRoutes(router, db)
	routes.CategoryRoutes(router, db)
	routes.ProductImageRoutes(router, db)
	routes.ProductRoutes(router, db)

	// Set up seeders
	// seeders.DBSeed(db)

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello World!",
		})
	})
	router.Static("/swagger-ui", "./static/swagger-ui")
	router.GET("/swagger.yaml", func(c *gin.Context) {
    c.File("./swagger.yaml")
	})

	log.Println("Server running on port " + port)
	if err := router.Run(":" + port); err != nil {
		log.Fatal("Failed to run server: ", err)
	}
}
