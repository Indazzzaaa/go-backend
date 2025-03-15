package auth_utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// Secret key for signing tokens ( change this for prod)
var jwtSecret = []byte("your-secret-key")

// GenerateToken creates a JWT token for given user ID.
func GenerateToken(userID uint) (string, error) {

	// create claims with expirations time
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	}

	// create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// sign and return token
	return token.SignedString(jwtSecret)

}

func GetJWTSecret() []byte {
	return jwtSecret
}
