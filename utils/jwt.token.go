package utils

import (
	"errors" 
	"github.com/golang-jwt/jwt/v5"
	"os"
	"time"
)

func GenerateToken(userID string) (string, error) {
	secretKey := os.Getenv("GO_JWT_SECRET")
	if secretKey == "" {
		return "", errors.New("JWT secret key not set")
	}

	expirationTime := time.Now().Add(1 * time.Hour)

	claims := &jwt.RegisteredClaims{
		ID:        userID,
		ExpiresAt: jwt.NewNumericDate(expirationTime),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(secretKey))
}
