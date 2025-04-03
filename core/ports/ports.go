package ports

import (
	"gin-starter/core/domain"
	"gin-starter/core/dto"
	"mime/multipart"

	"github.com/gin-gonic/gin"
)

type AuthService interface {
	SignUp(reqeust *dto.SignUpRequest) (*dto.TokenResponse, error)
	Login(request *dto.LoginRequest) (*dto.TokenResponse, error)
}

type AuthRepository interface {
	CreateUser(user *domain.User) (*domain.User, error)
	FindByEmail(email string) (*domain.User, error)
	ExistByEmail(email string) bool
}

type UserService interface {
	GetAccount(userId uint) (*dto.UserInfo, error)
	UpdateAvatar(avatarUrl string, userId uint) error
}

type UserRepository interface {
	FindById(userId uint) (*domain.User, error)
	UpdateAvater(user *domain.User) error
	FindFileByUrl(url string) (*domain.File, error)
	DeleteFileById(fileId uint) error
}

type FileService interface {
	UploadFile(file multipart.FileHeader, c *gin.Context) (*dto.UploadFileResponse, error)
	ServeFile(fileName string) (string, error)
}

type FileRepository interface {
	CreateFile(file *domain.File) (*domain.File, error)
	FindByFileName(filename string) (*domain.File, error)
	DeleteFile(id uint) error
}
