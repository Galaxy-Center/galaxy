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

// Default return default obj.
func (r Uint64Range) Default() Uint64Range {
	r.left = uint64(0)
	r.right = uint64(time.Now().UnixNano())
	return r
}

// Set setter of left and right.
func (r Uint64Range) Set(left, right uint64) Uint64Range {
	if left > right {
		log.Panic("time range left {} should less than right {}", left, right)
	}
	r.left = left
	r.right = right
	return r
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

// InnerDetector extension detector for mysql#Json or PSql#Jsonb field type.
type InnerDetector struct {
	level int8
	field string
	key1  string
	key2  string
	value interface{}
}

// SetLevel1 setter self to level 1.
func (i InnerDetector) SetLevel1(field, key1 string, value interface{}) InnerDetector {
	if field == "" {
		panic("field should be not null")
	}
	if key1 == "" {
		panic("key1 should be not null")
	}
	if value == nil {
		panic("value should be not null")
	}
	i.level = 1
	i.field = field
	i.key1 = key1
	i.value = value
	return i
}

// SetLevel2 setter self to level 2.
func (i InnerDetector) SetLevel2(field, key1, key2 string, value interface{}) InnerDetector {
	if field == "" {
		panic("field should be not null")
	}
	if key1 == "" || key2 == "" {
		panic("key1 or key2 should be not null")
	}
	if value == nil {
		panic("value should be not null")
	}
	i.level = 2
	i.field = field
	i.key1 = key1
	i.key2 = key2
	i.value = value
	return i
}

// GetLevel getter of level.
func (i *InnerDetector) GetLevel() int8 {
	return i.level
}

// GetField getter of field.
func (i *InnerDetector) GetField() string {
	return i.field
}

// GetKey1 getter of key1.
func (i *InnerDetector) GetKey1() string {
	return i.key1
}

// GetKey2 getter of key2.
func (i *InnerDetector) GetKey2() string {
	return i.key2
}

// GetValue getter of value.
func (i *InnerDetector) GetValue() interface{} {
	return i.value
}

// Condition builder the model query limit conditions.
type Condition struct {
	order            string
	timeRange        Uint64Range
	exlcudeInactived bool
	attachment       Attachment
	jsonQueries      []InnerDetector
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

// GetJSONQueries getter of jsonQueries.
func (c *Condition) GetJSONQueries() []InnerDetector {
	return c.jsonQueries
}

// SetExcludeInactived setter of excludeInactived.
func (c *Condition) SetExcludeInactived(excludeInactived bool) {
	c.exlcudeInactived = excludeInactived
}

// IsExcludeInactived getter of excludeInactived.
func (c *Condition) IsExcludeInactived() bool {
	return c.exlcudeInactived
}

// SetAttachment setter of attachment.
func (c *Condition) SetAttachment(attachment Attachment) {
	c.attachment = attachment
}

// AddJSONQueries adder of jsonQueries.
func (c *Condition) AddJSONQueries(qs ...InnerDetector) {
	for _, q := range qs {
		c.jsonQueries = append(c.jsonQueries, q)
	}
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
