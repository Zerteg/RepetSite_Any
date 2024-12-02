package models

import (
	"database/sql"
	"errors"
	"gorm.io/gorm"
	"time"
)

// User - структура пользователя
// @Description User object
// @Param user body models.User true "User data"
// @Success 200 {object} models.User
// @Failure 400 {object} models.Error
type User struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	Username  string `gorm:"not null" json:"username"`
	Email     string `gorm:"unique;not null" json:"email"`
	Phone     string `gorm:"unique;size:12" json:"phone"`
	Password  string `gorm:"not null" json:"password"`
}

// LoginInput структура для входных данных при регистрации
type LoginInput struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// CreateUser создает нового пользователя в базе данных
func CreateUser(db *gorm.DB, user *User) error {
	result := db.Create(user)
	return result.Error
}

// GetUserByUsername находит пользователя по имени
func GetUserByUsername(db *gorm.DB, username string) (*User, error) {
	var user User
	result := db.Where("username = ?", username).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

// GetUserByEmail находит пользователя по email
func GetUserByEmail(db *gorm.DB, email string) (User, error) {
	var user User
	result := db.Where("email = ?", email).First(&user)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return User{}, sql.ErrNoRows
		}
		return User{}, result.Error
	}
	return user, nil
}

// GetUserByEmail находит пользователя по email
func GetUserByPhone(db *gorm.DB, phone string) (User, error) {
	var user User
	result := db.Where("phone = ?", phone).First(&user)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return User{}, sql.ErrNoRows
		}
		return User{}, result.Error
	}
	return user, nil
}

func GetUserProfile(db *gorm.DB, userID string) (User, error) {
	var user User
	result := db.Where("id = ?", userID).First(&user) // Поиск пользователя по ID
	if result.Error != nil {
		return User{}, result.Error
	}
	return user, nil
}
