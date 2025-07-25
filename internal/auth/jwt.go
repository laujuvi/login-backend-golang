package auth

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(userID uint, duration time.Duration) (string, error) {
	claims := jwt.MapClaims{
		"sub": userID,
		"exp": time.Now().Add(duration).Unix(),
		"iat": time.Now().Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

func GenerateRefreshToken(userID uint, duration time.Duration) (string, error) {
	claims := jwt.MapClaims{
		"sub":  userID,
		"exp":  time.Now().Add(duration).Unix(),
		"iat":  time.Now().Unix(),
		"type": "refresh",
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}
