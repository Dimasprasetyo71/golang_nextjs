package utils

import (
    "fmt"
    "math/rand"
    "time"

    "github.com/brianvoe/gofakeit/v7"
    "github.com/shopspring/decimal"
)

func init() {
    rand.New(rand.NewSource(time.Now().UnixNano()))		
	gofakeit.Seed(time.Now().UnixNano())
}

func RandomInt(min, max int) int {
    return rand.Intn(max-min) + min
}

func RandomFloat(min, max float64) float64 {
    return min + rand.Float64()*(max-min)
}

func RandomDecimal(min, max float64) decimal.Decimal {
    randomFloat := min + rand.Float64()*(max-min)
    return decimal.NewFromFloat(randomFloat)
}

func RandomString(length int) string {
    bytes := make([]byte, length)
    for i := 0; i < length; i++ {
        bytes[i] = byte(rand.Intn(90-65) + 65)
    }
    return string(bytes)
}

func RandomBool() bool {
    return rand.Intn(2) == 1
}

func RandomStreetAddress() string {
    return fmt.Sprintf("%s, %s", gofakeit.Street(), gofakeit.City())
}

func RandomSecondaryAddress() string {
    types := []string{"Apt", "Suite", "Unit"}
    addrType := types[rand.Intn(len(types))]
    number := rand.Intn(100) + 1
    return fmt.Sprintf("%s %d", addrType, number)
}

func RandomPostCode() string {
    return gofakeit.Zip()
}
