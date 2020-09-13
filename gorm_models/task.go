package models

import (
	db "github.com/galaxy-center/galaxy/lifecycle"
)

// Task is an object representing the database table.
type Task struct {
	ID                 uint64 `gorm:"primaryKey,autoIncrement" json:"id" toml:"id" yaml:"id"`
	Name               string `gorm:"column:name" json:"name" toml:"name" yaml:"name"`
	Code               string `gorm:"column:code" json:"code" toml:"code" yaml:"code"`
	Type               string `gorm:"column:type" json:"type" toml:"type" yaml:"type"`
	Status             string `gorm:"column:status" json:"status" toml:"status" yaml:"status"`
	ExpiredAt          uint64 `gorm:"column:expired_at" json:"expired_at" toml:"expired_at" yaml:"expired_at"`
	Cron               string `gorm:"column:cron" json:"cron,omitempty" toml:"cron" yaml:"cron,omitempty"`
	Timeout            int    `gorm:"column:timeout" json:"timeout" toml:"timeout" yaml:"timeout"`
	SchedulingType     string `gorm:"column:scheduling_type" json:"scheduling_type" toml:"scheduling_type" yaml:"scheduling_type"`
	SchedulingCategory string `gorm:"column:scheduling_category" json:"scheduling_category" toml:"scheduling_category" yaml:"scheduling_category"`
	Assess             string `gorm:"column:assess" json:"assess" toml:"assess" yaml:"assess"`
	Executor           string `gorm:"column:executor" json:"executor" toml:"executor" yaml:"executor"`
	Active             bool   `gorm:"column:active" json:"active" toml:"active" yaml:"active"`
	CreatedAt          uint64 `gorm:"autoCreateTime:nano" json:"created_at" toml:"created_at" yaml:"created_at"`
	CreatedBy          string `gorm:"column:created_by" json:"created_by,omitempty" toml:"created_by" yaml:"created_by,omitempty"`
	UpdatedAt          uint64 `gorm:"autoUpdateTime:nano" json:"updated_at" toml:"updated_at" yaml:"updated_at"`
	UpdatedBy          string `gorm:"column:updated_by" json:"updated_by,omitempty" toml:"updated_by" yaml:"updated_by,omitempty"`
}

type Tabler interface {
	TableName() string
}

// TableName 会将 User 的表名重写为 `profiles`
func (Task) TableName() string {
	return "task"
}

// Create a single Task to db by *gorm.DB
func Create(task *Task) error {
	db := db.GetDB()
	err := db.Create(task).Error
	return err
}
