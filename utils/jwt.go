package utils

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

// GenerateJWT generates a JWT token based on the provided username and secret key
func GenerateJWT(username string, secretKey []byte) (string, error) {
	// Define the JWT claims
	claims := jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(time.Hour * 24).Unix(), // Token expires in 24 hours
	}

	// Create a new token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token with the secret key
	return token.SignedString(secretKey)
}