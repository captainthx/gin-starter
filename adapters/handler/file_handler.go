package handler

import (
	"gin-starter/core/ports"
	"net/http"

	"github.com/gin-gonic/gin"
)

type fileHandler struct {
	sv ports.FileService
}

func NewFileHandler(service ports.FileService) *fileHandler {
	return &fileHandler{sv: service}
}

func (f *fileHandler) UploadFile(c *gin.Context) {

	file, err := c.FormFile("file")

	if err != nil {
		HandlerError(c, err)
		return
	}

	result, err := f.sv.UploadFile(*file, c)
	if err != nil {
		HandlerError(c, err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "File uploaded successfully",
		"result":  result,
	})

}

func (f *fileHandler) ServeFile(c *gin.Context) {
	fileName := c.Param("fileName")
	filePath, err := f.sv.ServerFile(fileName)
	if err != nil {
		HandlerError(c, err)
		return
	}
	// Serve the image file
	c.File(filePath)
}
