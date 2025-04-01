package service

import (
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
)

type fileService struct {
	repo ports.FileRepository
}

func NewFileService(repo ports.FileRepository) ports.FileService {
	return &fileService{repo: repo}
}

// UploadFile implements ports.FileService.
func (f *fileService) UploadFile(file multipart.FileHeader, c *gin.Context) (*dto.UploadFileResponse, error) {

	user := c.MustGet("user").(*domain.User)
	contentType := file.Header.Get("Content-Type")

	var fileExtension string
	switch contentType {
	case "image/jpeg":
		fileExtension = ".jpg"
	case "image/png":
		fileExtension = ".png"
	case "image/webp":
		fileExtension = ".webp"
	default:
		return nil, errs.NewBadRequestError("file type not supported")
	}
	originalName := uuid.New().String()
	filename := originalName + fileExtension
	filePath := filepath.Join(config.UploadPath, filename)

	err := c.SaveUploadedFile(&file, filePath)
	if err != nil {
		logs.Error(fmt.Sprintf("UploadFile-[block].(save file to disk fail!). file:%v error:%v", file, err.Error()))
		return nil, errs.NewBadRequestError("save file to disk fail!")
	}

	// สร้าง file url
	fileUrl := config.FullBaseUrl + config.ImageBasePath + filename

	fileModel := &domain.File{
		UserID:       user.ID,
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
		logs.Error(fmt.Sprintf("UploadFile-[error].(save file fail). file:%v error:%v", file, err.Error()))
		return nil, errs.NewBadRequestError("save file fail!")
	}

	response := &dto.UploadFileResponse{
		FileName: result.Name,
		FileUrl:  result.Url,
		Size:     result.Size,
	}

	return response, nil
}

// ServerFile implements ports.FileService.
func (f *fileService) ServerFile(fileName string) (string, error) {
	filePath := filepath.Join(config.UploadPath, fileName)

	if _, err := os.Stat(filePath); err != nil {
		if os.IsNotExist(err) {
			return "", errs.NewNotFoundError("file not found.")
		}
		return "", errs.NewBadRequestError("file not found.")
	}
	return filePath, nil
}
