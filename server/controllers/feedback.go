package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"server/middlewares"
	"server/models"
)

type FeedbackController struct {
	DB *gorm.DB
}

// HandleFeedback обработчик для маршрута "/feedback"

// FeedbackHandler - метод для получения обратной связи
// @Summary Submit feedback
// @Description This endpoint allows a user to submit feedback
// @Tags Users
// @Accept json
// @Produce json
// @Param feedback body models.Feedback true "Feedback data"
// @Success 200 {object} models.Feedback
// @Failure 400 {object} models.Error
// @Router /feedback [post]
func (ctrl *FeedbackController) FeedbackHandler(c *gin.Context) {
	// Структура для входных данных
	var input models.Feedback
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Проверяем, что все поля заполнены
	if input.Name == "" || input.Phone == "" || input.Role == "" || input.Service == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "All fields are required"})
		return
	}

	feedback := models.Feedback{
		Name:    input.Name,
		Role:    input.Role,
		Phone:   input.Phone,
		Service: input.Service,
		Date:    input.Date,
	}

	models.CreateFeedback(middlewares.DB, &feedback)

	fmt.Printf(input.Name, input.Phone, input.Date, input.Role, input.Service)

	// Успешный ответ
	c.JSON(http.StatusOK, gin.H{"message": "Feedback submitted successfully"})
}
