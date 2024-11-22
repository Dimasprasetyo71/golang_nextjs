package config

import (
	// "DimasPrasetyo/backend-next/models"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Server struct {
	DB *gorm.DB
}

// you can use profesional test load for env like example https://dasarpemrogramangolang.novalagung.com/C-best-practice-configuration-env-var.html
func LoadEnv() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
}

func ConnectDB() *gorm.DB {
    LoadEnv()
    dsn := os.Getenv("POSTGRES_URL")
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatal("Failed to connect to database:", err)
    }
    fmt.Println("Database connected successfully")

    return db
}

// func Migrate(db *gorm.DB) {
// 	db.AutoMigrate(
// 		&models.User{},
// 		&models.Address{},
// 		&models.Category{},
// 		&models.Product{},
// 		&models.ProductImage{},
// 		&models.Payment{},
// 	)
// }