package utils

import (
	"github.com/golang-jwt/jwt/v5"
	"time"
)

var jwtKey = []byte("Authorization") // замените на секретный ключ

type Claims struct {
	UserID uint   `json:"user_id"`
	Role   string `json:"role"` // Добавляем роль
	jwt.RegisteredClaims
}

// Создание JWT
func GenerateJWT(userID uint, role string) (string, error) {
	claims := Claims{
		UserID: userID,
		Role:   role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)), // Токен действует 24 часа
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}

// Проверка JWT
func ValidateJWT(tokenString string) (uint, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil || !token.Valid {
		return 0, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return 0, err
	}

	// Извлечение userID как uint
	userID, ok := claims["userID"].(float64)
	if !ok {
		return 0, err
	}

	return uint(userID), nil
}
