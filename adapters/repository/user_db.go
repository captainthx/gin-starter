package repository

import (
	"gin-starter/core/domain"
	"gin-starter/core/ports"

	"gorm.io/gorm"
)

type userRepositoryDB struct {
	db *gorm.DB
}

func NewUserRepositoryDB(db *gorm.DB) ports.UserRepository {
	return &userRepositoryDB{db: db}
}

// FindById implements ports.UserRepository.
func (u *userRepositoryDB) FindById(userId uint) (*domain.User, error) {
	user := &domain.User{}
	err := u.db.Preload("Avatar").Where("id =?", userId).First(user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

// UpdateAvater implements ports.UserRepository.
func (u *userRepositoryDB) UpdateAvater(user *domain.User) error {
	err := u.db.Model(user).Update("avatar_id", user.AvatarID).Error
	if err != nil {
		return err
	}
	return nil
}

// DeleteFileByIdAndUserId implements ports.UserRepository.
func (u *userRepositoryDB) DeleteFileById(fileId uint) error {
	err := u.db.Where("id = ?", fileId).Delete(&domain.File{}).Error
	if err != nil {
		return err
	}
	return nil
}

// FindFileByUrlAndUserId implements ports.UserRepository.
func (u *userRepositoryDB) FindFileByUrl(url string) (*domain.File, error) {
	file := &domain.File{}
	err := u.db.Where("url =?", url).First(file).Error
	if err != nil {
		return nil, err
	}
	return file, nil
}
