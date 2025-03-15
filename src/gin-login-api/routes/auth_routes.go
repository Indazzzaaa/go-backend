package routes

import (
	auth_database "go-backend/src/gin-login-api/database"
	auth_middleware "go-backend/src/gin-login-api/middleware"
	auth_models "go-backend/src/gin-login-api/models"
	auth_utils "go-backend/src/gin-login-api/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// AuthRoutes sets up authentication routes
func AuthRoutes(router *gin.Engine) {

	authGroup := router.Group("/api")

	authGroup.POST("/register", RegisterUser)

	authGroup.GET("/login", LoginUser)

	authGroup.GET("/profile", auth_middleware.AuthMiddleware(), GetProfile)
	authGroup.POST("/logout", LogoutUser)

}

// Register User handles user registration
func RegisterUser(c *gin.Context) {

	var input auth_models.User

	//Bind JSON input
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return
	}

	// Hash the password
	if err := input.HashPassword(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}

	// save user in database
	if err := auth_database.DB.Create(&input).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	// Return success response
	c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully!"})

}

// Implement Login
// Accept Email and password
// Verifying user credentials
// Generating a JWT token if login is successful.
// LoginUser handles user authentication and JWT generation
func LoginUser(c *gin.Context) {
	var input auth_models.User

	// Bind JSON input
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Find user by email
	var user auth_models.User
	if err := auth_database.DB.Where("email = ?", input.Email).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	// Compare hashed password with input password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	// Generate JWT token
	token, err := auth_utils.GenerateToken(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	// set token as HTTP-only cookie
	c.SetCookie("jwt", token, 3600*24, "/", "", false, true)

	// Return token in response
	c.JSON(http.StatusOK, gin.H{
		"message": "Login successful!",
		"token":   token,
	})
}

func LogoutUser(c *gin.Context) {

	// clear the JWT cookie by seetting its max age to -1
	c.SetCookie("jwt", "", -1, "/", "", false, true)

	c.JSON(http.StatusOK, gin.H{
		"message": "Logout successful!",
	})

}

func GetProfile(c *gin.Context) {
	// Retrieve user object from context
	user, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	// Convert user to the correct type
	userData := user.(auth_models.User)

	// Return user profile (excluding password)
	c.JSON(http.StatusOK, gin.H{
		"id":       userData.ID,
		"username": userData.Username,
		"email":    userData.Email,
		"created":  userData.CreatedAt,
	})
}
