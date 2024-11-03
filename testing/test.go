package test

import (
	"DimasPrasetyo/backend-next/config"
	"DimasPrasetyo/backend-next/models"
	"testing"
)

func TestDB(t *testing.T) {
	config.ConnectDB()
	var product models.Product
	config.ConnectDB().Find(&product)
}

func TestDBMigrate(t *testing.T) {
	config.ConnectDB()
}

func BenchmarkDB(b *testing.B) {
	config.ConnectDB()
	for i := 0; i < b.N; i++ {
		config.ConnectDB()
	}
}
func Allmodels() {
	config.ConnectDB()

	var product models.Product
	config.ConnectDB().Find(&product)

	var user models.User
	config.ConnectDB().Find(&user)

	var category models.Category
	config.ConnectDB().Find(&category)

	var productImage models.ProductImage
	config.ConnectDB().Find(&productImage)

	var address models.Address
	config.ConnectDB().Find(&address)

}
