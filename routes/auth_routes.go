package routes

import (
	"gin-starter/adapters/handler"
	"gin-starter/adapters/repository"
	"gin-starter/core/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisAuthRoutes(router *gin.Engine, db *gorm.DB) {
	authRepo := repository.NewAuthRepositoryDB(db)
	authService := service.NewAuthService(authRepo)
	authHandler := handler.NewAuthHandler(authService)
	auth := router.Group("/v1/auth")
	{
		auth.POST("/sign-up", authHandler.SignUp)
		auth.POST("/login", authHandler.Login)
	}
}
