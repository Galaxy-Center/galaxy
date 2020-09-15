package models

import (
	"time"

	galaxyDB "github.com/galaxy-center/galaxy/lifecycle"
	"gorm.io/gorm"
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
	Actived            bool   `gorm:"column:actived" json:"active" toml:"active" yaml:"active"`
	CreatedAt          uint64 `gorm:"autoCreateTime:nano" json:"created_at" toml:"created_at" yaml:"created_at"`
	CreatedBy          string `gorm:"column:created_by" json:"created_by,omitempty" toml:"created_by" yaml:"created_by,omitempty"`
	UpdatedAt          uint64 `gorm:"autoUpdateTime:nano" json:"updated_at" toml:"updated_at" yaml:"updated_at"`
	UpdatedBy          string `gorm:"column:updated_by" json:"updated_by,omitempty" toml:"updated_by" yaml:"updated_by,omitempty"`
}

var TaskColumns = struct {
	ID                 string
	Name               string
	Code               string
	Type               string
	Status             string
	ExpiredAt          string
	Cron               string
	Timeout            string
	SchedulingType     string
	SchedulingCategory string
	Assess             string
	Executor           string
	Active             string
	CreatedAt          string
	CreatedBy          string
	UpdatedAt          string
	UpdatedBy          string
}{
	ID:                 "id",
	Name:               "name",
	Code:               "code",
	Type:               "type",
	Status:             "status",
	ExpiredAt:          "expired_at",
	Cron:               "cron",
	Timeout:            "timeout",
	SchedulingType:     "scheduling_type",
	SchedulingCategory: "scheduling_category",
	Assess:             "assess",
	Executor:           "executor",
	Active:             "active",
	CreatedAt:          "created_at",
	CreatedBy:          "created_by",
	UpdatedAt:          "updated_at",
	UpdatedBy:          "updated_by",
}

// Tabler defines the table name.
type Tabler interface {
	TableName() string
}

// TableName 会将 User 的表名重写为 `profiles`
func (Task) TableName() string {
	return "task"
}

// BeforeCreate do somethings, e.g. checks for the necessary fileds.
func (t *Task) BeforeCreate(tx *gorm.DB) (err error) {
	currTime := uint64(time.Now().UnixNano())
	if t.CreatedAt <= 0 {
		t.CreatedAt = currTime
	}
	if t.UpdatedAt <= 0 {
		t.UpdatedAt = currTime
	}
	return nil
}

// AfterCreate do somethings, e.g. debug log.
func (t *Task) AfterCreate(tx *gorm.DB) (err error) {
	// nothing doing
	return
}

// Create a single Task to db by *gorm.DB
func Create(task *Task) error {
	db := galaxyDB.GetDB()
	err := db.Create(task).Error
	return err
}

// BeforeUpdate do somethings, e.g. updating the updated_at value.
func (t *Task) BeforeUpdate(tx *gorm.DB) (err error) {
	t.UpdatedAt = uint64(time.Now().UnixNano())
	return
}

// AfterUpdate do somethings, e.g. update other database in the same transcation.
// 在同一个事务中更新数据
func (t *Task) AfterUpdate(tx *gorm.DB) (err error) {
	// nothing doing
	return
}

// Save a single task will be stored to db
// Note: all the fields will be updated to db, includes default value.
func Save(task *Task) error {
	db := galaxyDB.GetDB()
	err := db.Save(task).Error
	return err
}

// Updates updates from specific task that will not updating the zero value
// fileds to db.
// 只能保存非零字段
func Updates(task *Task) error {
	db := galaxyDB.GetDB()
	err := db.Model(task).Updates(task).Error
	return err
}

// UpdatesFromMap updates from specific task that will not updating the zero value
// fileds to db.
// 只能保存非零字段
func UpdatesFromMap(id uint64, values map[string]interface{}) error {
	db := galaxyDB.GetDB()
	err := db.Model(&Task{}).Where("id = ?", id).Updates(values).Error
	return err
}

// Delete delete permanently. 永久删除
func Delete(id uint64) error {
	db := galaxyDB.GetDB()
	err := db.Where(&Task{}, id).Error
	return err
}

// DeleteByActived delete softly. 软删除
func DeleteByActived(id uint64) error {
	db := galaxyDB.GetDB()
	err := db.Model(&Task{}).Where("id = ?", id).Update("actived", false).Error
	return err
}

// Get returns the task by specific id.
func Get(id uint64) (*Task, error) {
	db := galaxyDB.GetDB()
	var task Task
	err := db.First(&task, id).Error
	if err != nil {
		return &task, nil
	}
	return nil, err
}
