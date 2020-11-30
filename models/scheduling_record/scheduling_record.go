package schedulingrecord

import (
	"time"

	galaxyDB "github.com/galaxy-center/galaxy/lifecycle"
	models "github.com/galaxy-center/galaxy/models"
	"gorm.io/gorm"
)

// Status new->runnable->running->finished/failed
type Status string

const (
	// NEW the record initialized status.
	NEW Status = "NEW"
	// RUNNABLE NEW->RUNNABLE / RUNNING->RUNNABLE.
	RUNNABLE = "RUNNABLE"
	// RUNNING RUNNABLE->RUNNING / RUNNING->RUNNABLE/ RUNNING -> FINISHED / RUNNING->FAILED.
	RUNNING = "RUNNING"
	// FINISHED RUNNING->FINISHED.
	FINISHED = "FINISHED"
	// FAILED RUNNING->FAILED.
	FAILED = "FAILED"
)

// SchedulingRecord is an object representing the database table.
type SchedulingRecord struct {
	ID        uint64 `gorm:"primaryKey,autoIncrement" json:"id" toml:"id" yaml:"id"`
	TaskID    uint64 `gorm:"column:task_id" json:"task_id" toml:"task_id" yaml:"task_id"`
	Status    Status `gorm:"column:status" json:"status" toml:"status" yaml:"status"`
	Message   string `gorm:"column:message" json:"message" toml:"message" yaml:"message"`
	DeletedAt uint64 `gorm:"column:deleted_at" json:"deleted_at" toml:"deleted_at" yaml:"deleted_at"`
	CreatedAt uint64 `gorm:"autoCreateTime:nano" json:"created_at" toml:"created_at" yaml:"created_at"`
	CreatedBy string `gorm:"column:created_by" json:"created_by,omitempty" toml:"created_by" yaml:"created_by,omitempty"`
	UpdatedAt uint64 `gorm:"autoUpdateTime:nano" json:"updated_at" toml:"updated_at" yaml:"updated_at"`
	UpdatedBy string `gorm:"column:updated_by" json:"updated_by,omitempty" toml:"updated_by" yaml:"updated_by,omitempty"`
}

// SchedulingRecordColumns table field name.
var SchedulingRecordColumns = struct {
	ID        string
	TaskID    string
	Status    string
	Message   string
	DeletedAt string
	CreatedAt string
	CreatedBy string
	UpdatedAt string
	UpdatedBy string
}{
	ID:        "id",
	TaskID:    "task_id",
	Status:    "status",
	Message:   "message",
	DeletedAt: "deleted_at",
	CreatedAt: "created_at",
	CreatedBy: "created_by",
	UpdatedAt: "updated_at",
	UpdatedBy: "updated_by",
}

// Tabler defines the table name.
type Tabler interface {
	TableName() string
}

// TableName 会将 SchedulingRecord 的表名重写为 `scheduling_records`
func (SchedulingRecord) TableName() string {
	return "scheduling_records"
}

// AfterCreate do somethings, e.g. debug log.
func (r *SchedulingRecord) AfterCreate(tx *gorm.DB) (err error) {
	// nothing doing
	return
}

// Create a single SchedulingRecord to db by *gorm.DBs
func Create(record *SchedulingRecord) error {
	db := galaxyDB.GetDB()
	err := db.Create(record).Error
	return err
}

// BeforeUpdate do somethings, e.g. updating the updated_at value.
func (r *SchedulingRecord) BeforeUpdate(tx *gorm.DB) (err error) {
	r.UpdatedAt = uint64(time.Now().UnixNano())
	return
}

// AfterUpdate do somethings, e.g. update other database in the same transcation.
// 在同一个事务中更新数据
func (r *SchedulingRecord) AfterUpdate(tx *gorm.DB) (err error) {
	// nothing doing
	return
}

// Save a single task will be stored to db
// Note: all the fields will be updated to db, includes default value.
func Save(record *SchedulingRecord) error {
	db := galaxyDB.GetDB()
	err := db.Save(record).Error
	return err
}

// Updates updates from specific task that will not updating the zero value
// fields to db.
// 只能保存非零字段
func Updates(record *SchedulingRecord) error {
	db := galaxyDB.GetDB()
	err := db.Model(record).Updates(record).Error
	return err
}

// UpdatesFromMap updates from specific task that will not updating the zero value
// fields to db.
// 只能保存map包含字段
func UpdatesFromMap(id uint64, values map[string]interface{}) error {
	db := galaxyDB.GetDB()
	err := db.Model(&SchedulingRecord{}).Where("id = ?", id).Updates(values).Error
	return err
}

// Delete delete permanently. 永久删除
func Delete(id uint64) error {
	db := galaxyDB.GetDB()
	err := db.Delete(&SchedulingRecord{}, id).Error
	return err
}

// DeleteAt delete softly. 软删除
func DeleteAt(id uint64) error {
	db := galaxyDB.GetDB()
	err := db.Model(&SchedulingRecord{}).Where("id = ?", id).Update("deleted_at", time.Now().UnixNano()).Error
	return err
}

// Get returns the task by specific id.
func Get(id uint64) (*SchedulingRecord, error) {
	db := galaxyDB.GetDB()
	var record SchedulingRecord
	if err := db.First(&record, id).Error; err != nil {
		return nil, err
	}
	return &record, nil
}

// GetExcludeDeleted returns the task that excludes inactived by specific id.
func GetExcludeDeleted(id uint64) (*SchedulingRecord, error) {
	db := galaxyDB.GetDB()
	var record SchedulingRecord
	if err := db.Where("id = ?", id).Where("deleted_at = ?", 0).First(&record).Error; err != nil {
		return nil, err
	}
	return &record, nil
}

// PaginateQuery todo
func PaginateQuery(p *models.Pagination) (models.Response, error) {
	var response models.Response
	response.Page = p.GetPage()

	db := galaxyDB.GetDB()

	var total int64
	attached := models.Attach(p.BuildCondition())

	db.Model(&SchedulingRecord{}).Scopes(attached).Count(&total)
	response.Total = int(total)
	response.TotalPage = int(total)/p.GetPageSize() + 1

	var records []SchedulingRecord
	db.Scopes(attached, models.Paginate(p)).Find(&records)
	response.Data = records

	return response, nil
}
