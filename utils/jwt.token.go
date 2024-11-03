package utils

import (
	"errors" // Import package errors
	"github.com/golang-jwt/jwt/v5"
	"os"
	"time"
)

// GenerateToken generates a JWT token for a user with a given userID
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

	// Sign the token with the secret key
	return token.SignedString([]byte(secretKey))
}
