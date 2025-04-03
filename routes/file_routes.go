package routes

import (
	"gin-starter/adapters/handler"
	"gin-starter/adapters/repository"
	"gin-starter/core/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisFileRoutes(router *gin.Engine, db *gorm.DB) {
	fileRepo := repository.NewFileRepositoryDB(db)
	fileService := service.NewFileService(fileRepo)
	fileHandler := handler.NewFileHandler(fileService)

	file := router.Group("/v1/file")
	{
		file.POST("/upload", fileHandler.UploadFile)
		file.GET("/serve/:filename", fileHandler.ServeFile)
	}

}
