package utils

import (
	"github.com/golang-jwt/jwt/v5"
)

var JwtSecret []byte

type Claims struct {
	Username string `json:"username"`
	UserType string `json:"userType"`
	jwt.MapClaims
}

// GenerateJWT generates a new JWT token
func GenerateJWT(username, userType string) (string, error) {
	claims := &Claims{
		Username: username,
		UserType: userType,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(JwtSecret)

	return tokenString, err
}

// ValidateJWT validates the JWT token
func ValidateJWT(tokenString string) (*Claims, error) {
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return JwtSecret, nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, err
	}

	return claims, nil
}
