package controllers

import (
	"context"
	"net/http"
	"os"
	"time"

	"DimasPrasetyo/backend-next/models"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
	"golang.org/x/oauth2/google"
	"gorm.io/gorm"
)

type RegisterRequest struct {
	FirstName string `json:"first_name" validate:"required"`
	LastName  string `json:"last_name" validate:"required"`
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required"`
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}


var (
	githubOAuthConfig = &oauth2.Config{
		ClientID:     os.Getenv("GITHUB_CLIENT_ID"),
		ClientSecret: os.Getenv("GITHUB_CLIENT_SECRET"),
		RedirectURL:  "http://localhost:3000/auth/github/callback",
		Scopes:       []string{"user:email"},
		Endpoint:     github.Endpoint,
	}
)
var (
	googleOAuthConfig = &oauth2.Config{
		ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
		ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
		RedirectURL:  "http://localhost:3000/auth/callback/google",
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.email",
			"https://www.googleapis.com/auth/userinfo.profile",
		}, Endpoint: google.Endpoint,
	}
)

// Handler to start the Google login
func LoginWithGoogle(c *gin.Context) {
	url := googleOAuthConfig.AuthCodeURL("state")
	c.Redirect(http.StatusTemporaryRedirect, url)

}

// Callback handler for Google login
func GoogleCallback(c *gin.Context) {
	state := c.Query("state")
	code := c.Query("code")

	// Verify the state parameter to prevent CSRF attacks
	if state != "state" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid state parameter."})
		return
	}

	// Exchange the received code for a token
	token, err := googleOAuthConfig.Exchange(context.Background(), code)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not exchange token. Please try again."})
		return
	}

	// Use the token to get user info
	client := googleOAuthConfig.Client(context.Background(), token)
	resp, err := client.Get("https://www.googleapis.com/oauth2/v3/userinfo")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not retrieve user info. Please try again."})
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get user info from Google."})
		return
	}

	// Process user info and create or log in the user in your system
	// Assume user info is processed here

	// Generate JWT token
	// (JWT generation logic goes here)

	c.JSON(http.StatusOK, gin.H{"message": "User authenticated successfully with Google!"})
}

// Handler to start the GitHub login
func LoginWithGithub(c *gin.Context) {
	url := githubOAuthConfig.AuthCodeURL("state")
	c.Redirect(http.StatusTemporaryRedirect, url)
}

// Callback handler for GitHub login
func GithubCallback(c *gin.Context) {
	state := c.Query("state")
	code := c.Query("code")

	// Verify the state parameter to prevent CSRF attacks
	if state != "state" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid state parameter."})
		return
	}

	// Exchange the received code for a token
	token, err := githubOAuthConfig.Exchange(context.Background(), code)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not exchange token. Please try again."})
		return
	}

	// Use the token to get user info
	client := githubOAuthConfig.Client(context.Background(), token)
	resp, err := client.Get("https://api.github.com/user")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not retrieve user info. Please try again."})
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get user info from GitHub."})
		return
	}

	// Process user info and create or log in the user in your system
	// Assume user info is processed here

	c.JSON(http.StatusOK, gin.H{"message": "User authenticated successfully with GitHub!"})
}

// Register handles new user registration
func Register(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		req := new(RegisterRequest)
		if err := c.ShouldBindJSON(req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
			return
		}

		// Hash the password
		user := models.User{
			FirstName: req.FirstName,
			LastName:  req.LastName,
			Email:     req.Email,
		}
		if err := user.HashPassword(req.Password); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not hash password"})
			return
		}

		// Save user to the database
		if err := db.Create(&user).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create user"})
			return
		}

		c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully"})
	}
}

// Login handles user login and token generation
func Login(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		req := new(LoginRequest)
		if err := c.ShouldBindJSON(req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
			return
		}

		// Find user by email
		user := models.User{}
		if err := models.GetUserByEmail(db, req.Email, &user); err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
			return
		}

		// Check password
		if err := user.CheckPassword(req.Password); err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
			return
		}

		// Generate JWT token
		secretKey := os.Getenv("GO_JWT_SECRET")
		if secretKey == "" {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Secret key not set"})
			return
		}
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"user_id": user.ID,
			"exp":     time.Now().Add(time.Hour * 24).Unix(),
		})
		tokenString, err := token.SignedString([]byte(secretKey))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate token"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"token": tokenString})
	}
}

// ProtectedEndpoint is an example of a route that requires authentication
func ProtectedEndpoint(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "You are authorized!"})
}

// Logout simply acknowledges the logout request
func Logout(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "User logged out successfully"})
}

// GetAllUsers retrieves all users from the database
func GetAllUsers(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var users []models.User
		if err := db.Find(&users).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not retrieve users"})
			return
		}
		c.JSON(http.StatusOK, users)
	}
}

// GetUserByID retrieves a single user by their ID
func GetUserByID(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		var user models.User
		if err := db.First(&user, "id = ?", id).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			return
		}
		c.JSON(http.StatusOK, user)
	}
}

// UpdateUser updates an existing user's information by ID
func UpdateUser(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		var user models.User
		if err := db.First(&user, "id = ?", id).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			return
		}

		// Bind request data
		req := new(models.User)
		if err := c.ShouldBindJSON(req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
			return
		}

		// Only update allowed fields
		user.FirstName = req.FirstName
		user.LastName = req.LastName
		user.Email = req.Email
		user.Image = req.Image

		// Update user in the database
		if err := db.Save(&user).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not update user"})
			return
		}

		c.JSON(http.StatusOK, user)
	}
}

// DeleteUser removes a user from the database
func DeleteUser(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		var user models.User
		if err := db.First(&user, "id = ?", id).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			return
		}
		if err := db.Delete(&user).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not delete user"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
	}
}

// GetUserByEmail retrieves a user by their email
func GetUserByEmail(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		email := c.Param("email")
		var user models.User
		if err := db.First(&user, "email = ?", email).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			return
		}
		c.JSON(http.StatusOK, user)
	}
}

// UpdateUserPassword updates a user's password
func UpdateUserPassword(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		var user models.User
		if err := db.First(&user, "id = ?", id).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			return
		}

		// Bind request data
		req := new(models.User)
		if err := c.ShouldBindJSON(req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
			return
		}

		// Hash the new password before saving
		if err := user.HashPassword(req.Password); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not hash password"})
			return
		}

		// Update user in the database
		if err := db.Save(&user).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not update user"})
			return
		}

		c.JSON(http.StatusOK, user)
	}
}
