package common

import (
	"gorm.io/gorm"
)

type Pagination struct {
	Limit     int         `json:"limit"`
	Page      int         `json:"page"`
	Sort      string      `json:"sort"`
	TotalRows int64       `json:"total_rows"`
	TotalPage int         `json:"total_page"`
	Items     interface{} `json:"item"`
}

func (p *Pagination) GetOffset() int {
	return (p.GetPage() - 1) * p.GetLimit()
}

func (p *Pagination) GetLimit() int {
	if p.Limit == 0 {
		p.Limit = 10
	}
	return p.Limit
}

func (p *Pagination) GetPage() int {
	if p.Page <= 0 {
		p.Page = 1
	}
	return p.Page
}

func (p *Pagination) GetSort() string {
	if p.Sort == "" {
		p.Sort = "id desc"
	}
	return p.Sort
}

func Paginate(p *Pagination) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Offset(p.GetOffset()).Limit(p.GetLimit()).Order(p.GetSort())
	}
}

func GetTotalCount(db *gorm.DB, model interface{}) (int64, error) {
	var count int64
	err := db.Model(model).Count(&count).Error
	return count, err
}
