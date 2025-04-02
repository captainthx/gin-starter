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

// @Summary     Upload file
// @Description Upload file
// @Tags        File
// @Accept      json
// @Produce     json
// @Param       file formData  file true "File Upload Request"
// @Success     200 {object} dto.UploadFileResponse
// @Failure     400 {object} errs.AppError
// @Router      /v1/file/upload [post]
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
	c.JSON(http.StatusCreated, result)
}

// @Summary     Get File
// @Description Get a file by file name
// @Tags        File
// @Accept      json
// @Produce     application/*
// @Param       fileName path string true "File name to serve"
// @Success     200 {file} binary "File served successfully"
// @Failure     400 {object} errs.AppError
// @Failure     404 {object} errs.AppError
// @Router      /v1/file/serve/{fileName} [get]
// @Security    BearerToken
func (f *fileHandler) ServeFile(c *gin.Context) {
	fileName := c.Param("fileName")

	filePath, err := f.sv.ServeFile(fileName)
	if err != nil {
		HandlerError(c, err)
		return
	}
	// Serve the image file
	c.File(filePath)
}
