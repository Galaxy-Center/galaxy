package models

import (
	"gorm.io/gorm"
)

// Pagination wrapper the info of paginate.
type Pagination struct {
	PageSize int
	Page     int
}

// Response wrapper the pagination result.
type Response struct {
	Page      int
	TotalPage int
	Total     int
	Data      interface{}
}

// PaginationWrapper abstract interface wrapper of pagination infos confition.
type PaginationWrapper interface {
	Pagination() *Pagination
	From() uint64
	To() uint64
}

// Paginate returns a func with paging infomation.
func Paginate(p *Pagination) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if p.Page == 0 {
			p.Page = 1
		}
		if p.PageSize <= 0 {
			p.PageSize = 10
		}
		if p.PageSize > 100 {
			p.PageSize = 100
		}

		offset := (p.Page - 1) * p.PageSize
		return db.Offset(offset).Limit(p.PageSize)
	}
}
