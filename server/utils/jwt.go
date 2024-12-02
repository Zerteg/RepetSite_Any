package utils

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

var jwtKey = []byte("Authorization") // замените на секретный ключ

// Создание JWT
func GenerateJWT(userID uint) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userID": userID,
		"exp":    time.Now().Add(3 * time.Hour).Unix(),
	})
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
