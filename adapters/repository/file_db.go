package repository

import (
	"gin-starter/core/domain"
	"gin-starter/core/ports"

	"gorm.io/gorm"
)

type fileRepositoryDB struct {
	db *gorm.DB
}

func NewFileRepositoryDB(db *gorm.DB) ports.FileRepository {
	return &fileRepositoryDB{db: db}
}

// CreateFile implements ports.FileRepository.
func (f *fileRepositoryDB) CreateFile(file *domain.File) (*domain.File, error) {
	if err := f.db.Create(&file).Error; err != nil {
		return nil, err
	}
	return file, nil
}

// FindByFileName implements ports.FileRepository.
func (f *fileRepositoryDB) FindByFileName(filename string) (*domain.File, error) {
	file := domain.File{}
	err := f.db.Where("name =?", filename).First(&file).Error
	if err != nil {
		return nil, err
	}
	return &file, nil
}

// DeleteFile implements ports.FileRepository.
func (f *fileRepositoryDB) DeleteFile(id uint) error {
	err := f.db.Delete(&domain.File{}, id).Error
	if err != nil {
		return err
	}
	return nil
}
