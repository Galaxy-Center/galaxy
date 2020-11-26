package models

import (
	"time"

	log "github.com/sirupsen/logrus"
)

// Uint64Range wrapper range condition.
type Uint64Range struct {
	left  uint64
	right uint64
}

// NewUint64Range setter of left and right.
func NewUint64Range(left, right uint64) Uint64Range {
	if left > right {
		log.Panic("time range left {} should less than right {}", left, right)
	}
	return Uint64Range{
		left:  left,
		right: right,
	}
}

// GetLeft gtter of left.
func (r *Uint64Range) GetLeft() uint64 {
	if r.left < 0 {
		return 0
	}
	return r.left
}

// GetRight getter of right.
func (r *Uint64Range) GetRight() uint64 {
	if r.right <= 0 {
		return uint64(time.Now().UnixNano())
	}
	return r.right
}

// Condition builder the model query limit conditions.
type Condition struct {
	order            string
	timeRange        Uint64Range
	exlcudeInactived bool
	attachment       Attachment
}

// GetAttachmentOrNil returns (interface{}, bool) pair of key.
func (c *Condition) GetAttachmentOrNil(k string) (interface{}, bool) {
	if v, ok := c.attachment[k]; ok {
		return v, true
	}
	return nil, false
}

// SetTimeRange setter includes from, to.
func (c *Condition) SetTimeRange(r Uint64Range) {
	c.timeRange = r
}

// GetFrom getter of from.
func (c *Condition) GetFrom() uint64 {
	return c.timeRange.GetLeft()
}

//GetTo getter of to.
func (c *Condition) GetTo() uint64 {
	return c.timeRange.GetRight()
}

// SetExcludeInactived setter of excludeInactived.
func (c *Condition) SetExcludeInactived(excludeInactived bool) {
	c.exlcudeInactived = excludeInactived
}

// IsExcludeInactived getter of excludeInactived.
func (c *Condition) IsExcludeInactived() bool {
	return c.exlcudeInactived
}

// Collection wrapper slice condition.
type Collection struct {
	Values []interface{}
}

// Attachment alias of map[string]interface{}
type Attachment map[string]interface{}

// PaginationColumns defines all columns var for orm pagination.
var PaginationColumns = struct {
	Deleted   string
	TimeRange string
}{
	Deleted:   "excludeInactived",
	TimeRange: "timeRange",
}

