package models

import (
	"time"

	galaxyDB "github.com/galaxy-center/galaxy/lifecycle"
	"gorm.io/gorm"
)

// Type defines the type of task.
type Type int

const (
	// DelayJob default tasks, usually have a specific excution cycle.
	DelayJob Type = iota
	// DelayQueue queued tasks, this kind of task is usually one-time.
	DelayQueue
)

// Status defines by enum.
type Status int

const (
	// PENDING initialization status.
	PENDING Status = iota
	// ENABLED wirking status.
	ENABLED
	// DISABLED discarded status.
	DISABLED
)

// SchedulingCategory categories.
type SchedulingCategory int

const (
	// SINGLETON scheduling one node.
	SINGLETON SchedulingCategory = iota
	// MULTIPLE scheduling more nodes.
	MULTIPLE
)

// Executor defines.
type Executor int

const (
	// KAFKA from kafka
	KAFKA Executor = iota
	// RPC from rpc service
	RPC
	// HTTP from http app
	HTTP
)

// Task is an object representing the database table.
type Task struct {
	ID                 uint64             `gorm:"primaryKey,autoIncrement" json:"id" toml:"id" yaml:"id"`
	Name               string             `gorm:"column:name" json:"name" toml:"name" yaml:"name"`
	Code               string             `gorm:"column:code" json:"code" toml:"code" yaml:"code"`
	Type               Type               `gorm:"column:type" json:"type" toml:"type" yaml:"type"`
	Status             Status             `gorm:"column:status" json:"status" toml:"status" yaml:"status"`
	ExpiredAt          uint64             `gorm:"column:expired_at" json:"expired_at" toml:"expired_at" yaml:"expired_at"`
	Cron               string             `gorm:"column:cron" json:"cron,omitempty" toml:"cron" yaml:"cron,omitempty"`
	Timeout            int                `gorm:"column:timeout" json:"timeout" toml:"timeout" yaml:"timeout"`
	SchedulingCategory SchedulingCategory `gorm:"column:scheduling_category" json:"scheduling_category" toml:"scheduling_category" yaml:"scheduling_category"`
	Executor           Executor           `gorm:"column:executor" json:"executor" toml:"executor" yaml:"executor"`
	DeletedAt          uint64             `gorm:"column:deleted_at" json:"deleted_at" toml:"deleted_at" yaml:"deleted_at"`
	CreatedAt          uint64             `gorm:"autoCreateTime:nano" json:"created_at" toml:"created_at" yaml:"created_at"`
	CreatedBy          string             `gorm:"column:created_by" json:"created_by,omitempty" toml:"created_by" yaml:"created_by,omitempty"`
	UpdatedAt          uint64             `gorm:"autoUpdateTime:nano" json:"updated_at" toml:"updated_at" yaml:"updated_at"`
	UpdatedBy          string             `gorm:"column:updated_by" json:"updated_by,omitempty" toml:"updated_by" yaml:"updated_by,omitempty"`
}

// TaskColumns table field name.
var TaskColumns = struct {
	ID                 string
	Name               string
	Code               string
	Type               string
	Status             string
	ExpiredAt          string
	Cron               string
	Timeout            string
	SchedulingCategory string
	Executor           string
	DeletedAt          string
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
	SchedulingCategory: "scheduling_category",
	Executor:           "executor",
	DeletedAt:          "deleted_at",
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
	return "tasks"
}

// BeforeCreate do somethings, e.g. checks for the necessary fields.
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
// fields to db.
// 只能保存非零字段
func Updates(task *Task) error {
	db := galaxyDB.GetDB()
	err := db.Model(task).Updates(task).Error
	return err
}

// UpdatesFromMap updates from specific task that will not updating the zero value
// fields to db.
// 只能保存map包含字段
func UpdatesFromMap(id uint64, values map[string]interface{}) error {
	db := galaxyDB.GetDB()
	err := db.Model(&Task{}).Where("id = ?", id).Updates(values).Error
	return err
}

// Delete delete permanently. 永久删除
func Delete(id uint64) error {
	db := galaxyDB.GetDB()
	err := db.Delete(&Task{}, id).Error
	return err
}

// DeleteAt delete softly. 软删除
func DeleteAt(id uint64) error {
	db := galaxyDB.GetDB()
	err := db.Model(&Task{}).Where("id = ?", id).Update("deleted_at", time.Now().UnixNano()).Error
	return err
}

// Get returns the task by specific id.
func Get(id uint64) (*Task, error) {
	db := galaxyDB.GetDB()
	var task Task
	if err := db.First(&task, id).Error; err != nil {
		return nil, err
	}
	return &task, nil
}

// GetExcludeDeleted returns the task that excludes inactived by specific id.
func GetExcludeDeleted(id uint64) (*Task, error) {
	db := galaxyDB.GetDB()
	var task Task
	if err := db.Where("deleted_at = ?", 0).First(&task).Error; err != nil {
		return nil, err
	}
	return &task, nil
}

// TaskPagination implements of interface PaginationWrapper.
type TaskPagination struct {
	Page       Pagination
	Conditions map[string]interface{}
}

// Pagination returns the pagination of current query probe(查询探针).
func (tp *TaskPagination) Pagination() *Pagination {
	return &tp.Page
}

// Attachment returns attached info.
func (tp *TaskPagination) Attachment() Condition {
	var c Condition
	c.From = (tp.Page.Page - 1) * tp.Page.PageSize
	c.To = tp.Page.Page * tp.Page.PageSize
	if tp.Conditions["excludeInactived"] == true {
		c.ExlcudeInactived = true
		tp.Conditions["excludeInactived"] = nil
	}
	c.Attachment = tp.Conditions
	return c
}

// PaginateQuery todo
func PaginateQuery(pw PaginationWrapper) (Response, error) {
	var response Response
	response.Page = pw.Pagination().Page

	db := galaxyDB.GetDB()

	var total int64
	db.Scopes(Attach(pw.Attachment())).Count(&total)
	response.Total = int(total)
	response.TotalPage = int(total)/pw.Pagination().PageSize + 1

	var tasks []Task
	db.Scopes(Attach(pw.Attachment()), Paginate(pw.Pagination())).Find(&tasks)
	response.Data = tasks

	return response, nil
}
