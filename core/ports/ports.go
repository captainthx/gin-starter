package ports

import (
	"gin-starter/core/domain"
	"gin-starter/core/dto"
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
}

type UserRepository interface {
}

type FileService interface {
}

type FileRepository interface {
}
