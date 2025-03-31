package repository

import (
	"gin-starter/core/domain"
	"gin-starter/core/ports"

	"gorm.io/gorm"
)

type authRepositoryDB struct {
	db *gorm.DB
}

func NewAuthRepositoryDB(db *gorm.DB) ports.AuthRepository {
	return &authRepositoryDB{db: db}
}

// CreateUser implements ports.AuthRepository.
func (a *authRepositoryDB) CreateUser(user *domain.User) (*domain.User, error) {
	if err := a.db.Create(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

// ExistByEmail implements ports.AuthRepository.
func (a *authRepositoryDB) ExistByEmail(email string) bool {
	var exists bool
	query := `SELECT EXISTS(SELECT 1 FROM users WHERE email = ?)`
	err := a.db.Raw(query, email).Scan(&exists).Error
	if err != nil {
		return false
	}
	return exists
}

// FindByEmail implements ports.AuthRepository.
func (a *authRepositoryDB) FindByEmail(email string) (*domain.User, error) {
	user := domain.User{}
	err := a.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
