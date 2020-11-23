package models

import (
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

// Pagination wrapper the info of paginate.
type Pagination struct {
	PageSize int
	Page     int
}

// Condition wrapper the model query limit conditions.
type Condition struct {
	From             int
	To               int
	ExlcudeInactived bool
	Attachment       map[string]interface{}
}

// Collection wrapper slice condition.
type Collection struct {
	Values []interface{}
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
	Attachment() Condition
}

// Paginate returns a func with paging infomation.
// Ideally, the func includes: offset, limit, created_at, actived.
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

// Attach returns the attached db with input condition.
// Note: will attach fields by go interface assertion.
func Attach(c Condition) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		tx := db.Where("created_at BETWEEN ? AND ?", c.From, c.To)
		if c.ExlcudeInactived {
			tx = tx.Where("actived = ?", 0)
		}
		for k, v := range c.Attachment {
			switch t := v.(type) {
			case string, bool:
			case float32, float64, complex64, complex128:
			case int8, int16, int32, int, int64, uint8, uint16, uint32, uint, uint64, uintptr:
				tx = tx.Where(k+" = ?", t)
				continue
			case Collection:
				tx = tx.Where(k+"IN ?", t.Values)
			default:
				log.Error("Not support data type.")
			}
		}
		return tx
	}
}
