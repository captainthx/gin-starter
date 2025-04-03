package service

import (
	"fmt"
	errs "gin-starter/common/err"
	"gin-starter/core/dto"
	"gin-starter/core/ports"
	"gin-starter/logs"
	"os"
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
		Avartar: user.Avatar.Url,
	}

	return response, nil
}

// UpdateAvatar implements ports.UserService.
func (u *userService) UpdateAvatar(avatarUrl string, userId uint) error {

	user, err := u.repo.FindById(userId)
	if err != nil {
		logs.Warn(fmt.Sprintf("UpdateAvatar-[block].(user not found). userId:%v error:%v", userId, err.Error()))
		return errs.NewNotFoundError("user not found!")
	}

	file, err := u.repo.FindFileByUrl(avatarUrl)
	if err != nil {
		logs.Error(fmt.Sprintf("UpdateAvatar-[error].(file not found). error:%v", err.Error()))
		return errs.NewNotFoundError("file not found")
	}

	if file.ID != user.AvatarID {
		err := u.repo.DeleteFileById(user.AvatarID)
		if err != nil {
			logs.Error(fmt.Sprintf("UpdateAvatar-[error].(delete file in db error) error:%v", err.Error()))
			return err
		}
		err = os.Remove(file.Path)
		if err != nil {
			logs.Error(fmt.Sprintf("UpdateAvatar-[error].(deleta file form disk fail). error:%v", err.Error()))
			return err
		}
	}

	user.AvatarID = file.ID
	err = u.repo.UpdateAvater(user)
	if err != nil {
		logs.Error(fmt.Sprintf("UpdateAvatar-[error].(unknow error). error:%v", err.Error()))
		return err
	}

	return nil
}
