package service

import (
	"fmt"
	errs "gin-starter/common/err"
	"gin-starter/core/domain"
	"gin-starter/core/dto"
	"gin-starter/core/ports"
	"gin-starter/logs"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type authService struct {
	repo ports.AuthRepository
}

func NewAuthService(repo ports.AuthRepository) ports.AuthService {
	return &authService{repo: repo}

}

const (
	InvalidCredentialsMsg = "invalid credentials"
)

// Login implements ports.AuthService.
func (a *authService) Login(request *dto.LoginRequest) (*dto.TokenResponse, error) {

	result, err := a.repo.FindByEmail(request.Email)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errs.NewNotFoundError(InvalidCredentialsMsg)
		}
	}

	if result == nil { // ตรวจสอบว่า result ไม่ใช่ nil
		return nil, errs.NewNotFoundError(InvalidCredentialsMsg)
	}

	err = bcrypt.CompareHashAndPassword([]byte(result.Password), []byte(request.Password))

	if err != nil {
		logs.Error(fmt.Sprintf("Login-[block].(password not match). reqeust:%v error:%v ", request.Email, err.Error()))
		return nil, errs.NewBadRequestError(InvalidCredentialsMsg)
	}
	tokenResponse := GenerateToken(result.ID)
	return tokenResponse, nil
}

// SignUp implements ports.AuthService.
func (a *authService) SignUp(request *dto.SignUpRequest) (*dto.TokenResponse, error) {

	user := &domain.User{}

	hasPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		logs.Error(fmt.Sprintf("SignUp-[block].(error hasing password). request:%v error:%v", request.Email, err.Error()))
		return nil, errs.NewUnexpectedError("failed to hash password")
	}

	exist := a.repo.ExistByEmail(request.Email)
	if exist {
		logs.Warn(fmt.Sprintf("SingUp-[block].(email duplicate). request:%v ", request.Email))
		return nil, errs.NewBadRequestError("email already exists")
	}

	user.Password = string(hasPassword)
	user.Email = request.Email
	user, err = a.repo.CreateUser(user)
	if err != nil {
		logs.Error(fmt.Sprintf("SingUp-[error].(database error). request:%v error:%v ", request.Email, err.Error()))
		return nil, errs.NewUnexpectedError("unknow error")
	}

	tokenResponse := GenerateToken(user.ID)
	return tokenResponse, nil
}
