package models

import (
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

// Order type for order
type Order string

// Orders
const (
	ASC  Order = "ASC"
	DESC Order = "DESC"
)

const (
	defaultPage  = 1
	defaultLimit = 10
	defaultOrder = DESC
	maxLimit     = 100
)

// Pagination builder the info of paginate.
type Pagination struct {
	pageSize   int
	page       int
	attachment Attachment
}

// BuildCondition builder func of condition.
func (p *Pagination) BuildCondition() *Condition {
	var c Condition
	for k, v := range p.attachment {
		if k == PaginationColumns.TimeRange {
			c.SetTimeRange(v.(Uint64Range))
			delete(p.attachment, k)
			continue
		}
		if k == PaginationColumns.Deleted && v == true {
			c.SetExcludeInactived(true)
			delete(p.attachment, k)
			continue
		}
	}
	c.SetAttachment(p.attachment)
	return &c
}

// NewPagination returns default obj of Pagination.
func NewPagination() *Pagination {
	return &Pagination{
		pageSize: defaultLimit,
		page:     defaultPage,
	}
}

// SetPageSize setter of pageSize.
func (p *Pagination) SetPageSize(pageSize int) {
	if pageSize <= 0 {
		p.pageSize = defaultLimit
	} else {
		p.pageSize = pageSize
	}
}

// GetPageSize geeter of pageSize.
func (p *Pagination) GetPageSize() int {
	return p.pageSize
}

// SetPage setter of page.
func (p *Pagination) SetPage(page int) {
	if page < defaultPage {
		p.page = defaultPage
	} else {
		p.page = page
	}
}

// GetPage getter of page.
func (p *Pagination) GetPage() int {
	return p.page
}

// SetAttachment setter of attachment.
func (p *Pagination) SetAttachment(a Attachment) {
	p.attachment = a
}

// GetAttachment getter of attachment.
func (p *Pagination) GetAttachment() Attachment {
	return p.attachment
}

// Response wrapper the pagination result.
type Response struct {
	Page      int
	TotalPage int
	Total     int
	Data      interface{}
}

// Paginate returns a func with paging infomation.
// Ideally, the func includes: offset, limit, created_at, actived.
func Paginate(p *Pagination) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if p.GetPage() == 0 {
			p.SetPage(defaultPage)
		}
		if p.GetPageSize() <= 0 {
			p.SetPageSize(defaultLimit)
		}
		if p.GetPageSize() > maxLimit {
			p.SetPageSize(maxLimit)
		}
		offset := (p.GetPage() - 1) * p.GetPageSize()
		return db.Offset(offset).Limit(p.GetPageSize())
	}
}

// Attach returns the attached db with input condition.
// Note: will attach fields by go interface assertion.
func Attach(c *Condition) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		tx := db.Where("created_at BETWEEN ? AND ?", c.GetFrom(), c.GetTo())
		if c.IsExcludeInactived() {
			tx = tx.Where("deleted_at = ?", 0)
		}
		for k, v := range c.attachment {
			switch t := v.(type) {
			case string, bool,
				float32, float64, complex64, complex128,
				int8, int16, int32, int, int64, uint8, uint16, uint32, uint, uint64, uintptr:
				tx = tx.Where(k+" = ?", t)
				break
			case Collection:
				tx = tx.Where(k+" IN ?", t.Values)
				break
			case Uint64Range:
				tx = tx.Where(k+" BETWEEN ? AND ?", t.GetLeft(), t.GetRight()) // common is numberic field.
				break
			default:
				log.Error("Not support data type.")
			}
		}
		return tx
	}
}
