package service

import (
	"fmt"
	errs "gin-starter/common/err"
	"gin-starter/core/dto"
	"gin-starter/core/ports"
	"gin-starter/logs"
)

type userService struct {
	repo ports.UserRepository
}

func NewUserService(repo ports.UserRepository) ports.UserService {
	return &userService{repo: repo}
}

// GetAccount implements ports.UserService.
func (u *userService) GetAccount(userId uint) (*dto.UserInfo, error) {
	user, err := u.repo.FindById(userId)
	if err != nil {
		logs.Warn(fmt.Sprintf("GetAccount-[block].(user not found!). userId:%v error:%v", userId, err.Error()))
		return nil, errs.NewNotFoundError("user not found!")
	}

	response := &dto.UserInfo{
		Id:      user.ID,
		Email:   user.Email,
		Avartar: user.Avartar.Url,
	}

	return response, nil
}

// UpdateAvatar implements ports.UserService.
func (u *userService) UpdateAvatar(request *dto.UpdateAvartarRequest) error {
	panic("unimplemented")
}
