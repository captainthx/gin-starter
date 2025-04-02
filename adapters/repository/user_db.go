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
	err := u.db.Preload("Avartar").Where("id =?", userId).First(&user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

// UpdateAvater implements ports.UserRepository.
func (u *userRepositoryDB) UpdateAvater(user *domain.User) error {
	panic("unimplemented")
}
