package domain

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email    string `json:"email" gorm:"type:varchar(255);NOT NULL;unique"`
	Password string `json:"password" gorm:"type:varchar(255);NOT NULL"`
	AvatarID uint   `json:"avatar_id" gorm:"default:NULL"`
	Avatar   File   `gorm:"foreignKey:AvatarID;constraint:OnDelete:SET NULL;"`
}

type File struct {
	gorm.Model
	OriginalName string `json:"original_name" gorm:"type:varchar(255)"`
	Name         string `json:"name" gorm:"type:varchar(50);NOT NULL"`
	Path         string `json:"path" gorm:"type:varchar(255);NOT NULL"`
	Url          string `json:"url" gorm:"type:varchar(255);NOT NULL"`
	ContentType  string `json:"content_type" gorm:"type:varchar(22);NOT NULL"`
	Type         string `json:"type" gorm:"default:image;type:varchar(15);NOT NULL"`
	Extension    string `json:"extension" gorm:"type:varchar(10);NOT NULL"`
	Size         int64  `json:"size" gorm:"type:bigint;NOT NULL"`
}
