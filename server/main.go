package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"server/controllers"
	"server/middlewares"
	"server/models"
	"time"
)

var DB *gorm.DB

// @title Farengeit API
// @version 1.0
// @description API Server for Farengeit_shop Application
// @host localhost:8080
// @BasePath /

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {

	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:63342"}, // Укажите origin, который будет использоваться
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	var err error
	dsn := "user=postgres password=qwerty dbname=repet_any port=5433 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	// Миграция
	DB.AutoMigrate(&models.User{}, &models.Feedback{})
	middlewares.DB = DB // сохраняем подключение к базе данных в мидлваре

	feedbackController := &controllers.FeedbackController{
		DB: DB,
	}

	// Публичные маршруты
	router.POST("/register", controllers.RegisterUser)
	router.POST("/login", controllers.LoginUser)
	router.POST("/feedback", feedbackController.FeedbackHandler)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Группа защищенных маршрутов
	protectedRoutes := router.Group("/protected")
	protectedRoutes.Use(middlewares.AuthMiddleware()) // добавляем AuthMiddleware

	// Обработчик для статичных файлов (изображений)
	protectedRoutes.Static("/media", "./media")

	// Запуск сервера
	router.Run(":8080")
}
