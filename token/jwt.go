package token

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateAccessToken(email string) (string, error) {
	secret := os.Getenv("JWT_ACCESS_KEY")
	expirationSeconds := 120

	claims := jwt.MapClaims{
		"email": email,
		"exp":   time.Now().Add(time.Duration(expirationSeconds) * time.Second).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}

func GenerateRefreshToken(email string) (string, error) {
	secret := os.Getenv("JWT_REFRESH_KEY")
	expirationDays := 7

	claims := jwt.MapClaims{
		"email": email,
		"exp":   time.Now().Add(time.Duration(expirationDays*24) * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}

func ValidateToken(tokenString string, secretKey string) (*jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(secretKey), nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid token")
	}

	return &claims, nil
}
