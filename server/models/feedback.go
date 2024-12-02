package models

import (
	"fmt"
	"gorm.io/gorm"
	"time"
)

// Feedback - структура для хранения отзывов
// @Description Feedback object
type Feedback struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	Name      string    `gorm:"not null" json:"name"`
	Role      string    `gorm:"not null" json:"role"`
	Phone     string    `gorm:"not null" json:"phone"`
	Service   string    `gorm:"not null" json:"service"`
	Date      string    `gorm:"not null" json:"date"`
}

// CreateFeedback создает новый отзыв в базе данных
func CreateFeedback(db *gorm.DB, feedback *Feedback) error {
	// Преобразуем строку в тип time.Time
	date, err := time.Parse("2006-01-02", feedback.Date)
	if err != nil {
		return fmt.Errorf("invalid date format: %v", err)
	}

	// Обновляем поле Date в структуре
	feedback.Date = date.Format("02.01.2006")

	// Сохраняем отзыв в базе данных
	result := db.Create(feedback)
	return result.Error
}
