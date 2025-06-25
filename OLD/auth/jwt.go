package auth

import (
	"time"
	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(userID int, email string) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"email":   email,
		"exp":     time.Now().Add(time.Minute * time.Duration(TokenLifetimeMinutes)).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(SecretKey))
}
