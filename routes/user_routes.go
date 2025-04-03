package routes

import (
	"gin-starter/adapters/handler"
	"gin-starter/adapters/repository"
	"gin-starter/core/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisUserRoutes(router *gin.Engine, db *gorm.DB) {
	userRepo := repository.NewUserRepositoryDB(db)
	userService := service.NewUserService(userRepo)
	userHadler := handler.NewUserHandler(userService)

	user := router.Group("/v1/user")

	{
		user.GET("/me", userHadler.GetUser)
		user.POST("/update", userHadler.UpdataAvatar)
	}

}
