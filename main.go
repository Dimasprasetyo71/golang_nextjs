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

	router := gin.New()

	// Initialize services
	router.Use(cors.Default())
	router.Use(middleware.ErrorHandlerMiddleware())
	router.Use(middleware.RateLimitMiddleware())
	db := config.ConnectDB()
	config.InitMidtrans()
	gin.SetMode(gin.ReleaseMode)

	// Set up routes
	routes.UserRoutes(router, db)
	routes.AddressRoutes(router, db)
	routes.CategoryRoutes(router, db)
	routes.ProductImageRoutes(router, db)
	routes.ProductRoutes(router, db)
	routes.PaymentRoutes(router, db)
	// router.PaymentRoutes()
	// Set up seeders
	// seeders.DBSeed(db)

	router.NoRoute(NotFound)
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

func NotFound(c *gin.Context) {
	c.JSON(http.StatusNotFound, gin.H{"error": "Not Found"})
}
