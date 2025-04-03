package service

import (
	"errors"
	"fmt"
	errs "gin-starter/common/err"
	"gin-starter/config"
	"gin-starter/core/domain"
	"gin-starter/core/dto"
	"gin-starter/core/ports"
	"gin-starter/logs"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type fileService struct {
	repo ports.FileRepository
}

func NewFileService(repo ports.FileRepository) ports.FileService {
	return &fileService{repo: repo}
}

// UploadFile implements ports.FileService.
func (f *fileService) UploadFile(file multipart.FileHeader, c *gin.Context) (*dto.UploadFileResponse, error) {

	contentType := file.Header.Get("Content-Type")
	fmt.Println("content type", contentType)
	fileExtension := mapFileExtension(contentType)

	if !isAllowedContentType(contentType) {
		logs.Warn(fmt.Sprintf("UploadFile-[block].(file type not supported). file:%v contentType:%v", file.Filename, contentType))
		return nil, errs.NewBadRequestError("file type not supported")
	}

	originalName := uuid.New().String()
	filename := originalName + fileExtension
	filePath := filepath.Join(config.UploadPath, filename)

	err := c.SaveUploadedFile(&file, filePath)
	if err != nil {
		logs.Error(fmt.Sprintf("UploadFile-[error].(save file to disk fail!). file:%v error:%v", file, err.Error()))
		return nil, errs.NewBadRequestError("save file to disk fail!")
	}

	// สร้าง file url
	fileUrl := config.FullBaseUrl + config.ImageBasePath + filename

	fileModel := &domain.File{
		OriginalName: originalName,
		Name:         filename,
		ContentType:  contentType,
		Path:         filePath,
		Url:          fileUrl,
		Extension:    strings.Split(fileExtension, ".")[1],
		Size:         int64(file.Size),
	}

	result, err := f.repo.CreateFile(fileModel)
	if err != nil {
		logs.Error(fmt.Sprintf("UploadFile-[error].(save file fail). file:%v error:%v", fileModel, err.Error()))
		return nil, errs.NewBadRequestError("save file fail!")
	}

	response := &dto.UploadFileResponse{
		FileName: result.Name,
		FileUrl:  result.Url,
		Size:     result.Size,
	}

	return response, nil
}

// ServeFile implements ports.FileService.
func (f *fileService) ServeFile(fileName string) (string, error) {

	file, err := f.repo.FindByFileName(fileName)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		logs.Error(fmt.Sprintf("ServerFile-[block].(file not found). filename:%v error:%v", fileName, err.Error()))
		return "", errs.NewNotFoundError("file not found.")
	}

	if _, err := os.Stat(file.Path); err != nil {
		if os.IsNotExist(err) {
			return "", errs.NewNotFoundError("file not found.")
		}
	}
	return file.Path, nil
}

func isAllowedContentType(contentType string) bool {
	allowedType := []string{"image/jpeg", "image/png", "image/webp"}
	for _, allowed := range allowedType {
		if contentType == allowed {
			return true
		}
	}
	return false
}

func mapFileExtension(contentType string) string {
	fileExtension := ""
	switch contentType {
	case "image/jpeg":
		fileExtension = ".jpg"
	case "image/png":
		fileExtension = ".png"
	case "image/webp":
		fileExtension = ".webp"
	}
	return fileExtension
}
