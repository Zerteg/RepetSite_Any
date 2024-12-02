package rotes

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/gorm"
	"server/controllers"
)

var DB *gorm.DB

func RegisterAuthRoutes(router *gin.Engine) {

	// Создаем подгруппу маршрутов для аутентификации
	authRoutes := router.Group("/auth")

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Маршрут для регистрации
	authRoutes.POST("/register", controllers.RegisterUser)

	// Маршрут для входа
	authRoutes.POST("/login", controllers.LoginUser)

}
