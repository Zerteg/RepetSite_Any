package controllers

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"net/http"
	"server/middlewares"
	"server/models"
	"server/utils"
)

var db *gorm.DB

// RegisterUser - метод для регистрации пользователя
// @Summary User registration
// @Description This endpoint allows a new user to register by providing username, email, and password
// @Tags Users
// @Accept json
// @Produce json
// @Param user body models.User true "User data"
// @Success 200 {object} models.User
// @Failure 400 {object} models.Error
// @Router /register [post]
func RegisterUser(c *gin.Context) {
	var input models.User
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Проверяем, существует ли уже пользователь с таким email
	existingUser, err := models.GetUserByEmail(middlewares.DB, input.Email)
	if err != nil && err != sql.ErrNoRows {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при проверке email"})
		return
	}

	// Проверяем, существует ли пользователь с таким номером телефона
	existingPhone, err := models.GetUserByPhone(middlewares.DB, input.Phone)
	if err != nil && err != sql.ErrNoRows {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при проверке номера телефона"})
		return
	}

	// Если email уже существует
	if existingUser != (models.User{}) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Такой email уже существует"})
		return
	}

	// Если номер телефона уже существует
	if existingPhone != (models.User{}) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Такой номер телефона уже существует"})
		return
	}

	// Создаем нового пользователя
	password_hash, err := utils.HashPassword(input.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error hashing password"})
		return
	}

	user := models.User{
		Username: input.Username,
		Email:    input.Email,
		Phone:    input.Phone,
		Password: password_hash, // Хешированный пароль
		Role:     "user",
	}

	err = models.CreateUser(middlewares.DB, &user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User registered successfully"})
}

// Функция для входа пользователя

// LoginUser - метод для входа пользователя
// @Summary User login
// @Description This endpoint allows a user to log in with username and password
// @Tags Users
// @Accept json
// @Produce json
// @Param user body models.User true "Login credentials"
// @Success 200 {object} models.User
// @Failure 400 {object} models.Error
// @Router /login [post]
func LoginUser(c *gin.Context) {
	var input models.LoginInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Проверка существования пользователя
	user, err := models.GetUserByEmail(middlewares.DB, input.Email)

	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Неправильный email"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch user"})
		return
	}

	// Сравнение пароля
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Неправильный пароль"})
		return
	}

	// Создание и отправка JWT-токена
	token, err := utils.GenerateJWT(user.ID, user.Role) // Передаём роль
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": token,
		"role":  user.Role, // Добавляем роль в ответ
	})

}

func GetUserProfile(c *gin.Context) {
	userID := c.GetString("userID") // Получаем ID пользователя из контекста, который был сохранен в AuthMiddleware

	// Получаем профиль пользователя из базы данных
	var user models.User
	user, err := models.GetUserProfile(db, userID) // вызываем функцию для получения профиля по ID
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch user profile"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"profile": user})
}

func AdminOnlyEndpoint(c *gin.Context) {
	role := c.GetString("role")
	if role != "admin.html" {
		c.JSON(http.StatusForbidden, gin.H{"error": "Access denied"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Welcome, admin.html!"})
}
