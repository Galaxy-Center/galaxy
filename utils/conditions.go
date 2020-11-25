package utils

import (
	log "github.com/sirupsen/logrus"
)

// Uint64Range wrapper range condition.
type Uint64Range struct {
	left  uint64
	right uint64
}

// SetValues setter of left and right.
func (r *Uint64Range) SetValues(left, right uint64) {
	if left > right {
		log.Panic("time range left {} should less than right {}", left, right)
	}
	r.left = left
	r.right = right
}

// GetLeft gtter of left.
func (r *Uint64Range) GetLeft() uint64 {
	if r.GetLeft() < 0 {
		return 0
	}
	return r.GetLeft()
}

// GetRight getter of right.
func (r *Uint64Range) GetRight() uint64 {
	return r.right
}

// Condition builder the model query limit conditions.
type Condition struct {
	timeRange        Uint64Range
	exlcudeInactived bool
	attachment       Attachment
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
