package utils

import (
	"golang.org/x/crypto/bcrypt"
)

// HashPassword хеширует пароль перед сохранением в базе данных
func HashPassword(password string) (string, error) {
	password_hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(password_hash), nil
}
