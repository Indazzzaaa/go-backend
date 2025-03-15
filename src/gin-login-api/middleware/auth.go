package auth_middleware

import (
	auth_database "go-backend/src/gin-login-api/database"
	auth_models "go-backend/src/gin-login-api/models"
	auth_utils "go-backend/src/gin-login-api/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

// AuthMiddleware protects routes by verifying JWT tokens
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Retrieve the JWT token from the cookie
		token, err := c.Cookie("jwt")
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization token missing"})
			c.Abort()
			return
		}

		// Parse and validate the JWT token
		parsedToken, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
			return auth_utils.GetJWTSecret(), nil
		})
		if err != nil || !parsedToken.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
			c.Abort()
			return
		}

		// Extract user ID from token claims
		claims, ok := parsedToken.Claims.(jwt.MapClaims)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token claims"})
			c.Abort()
			return
		}

		userID := uint(claims["user_id"].(float64))

		// Fetch user from database
		var user auth_models.User
		if err := auth_database.DB.First(&user, userID).Error; err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
			c.Abort()
			return
		}

		// Attach user to context and continue
		c.Set("user", user)
		c.Next()
	}
}
