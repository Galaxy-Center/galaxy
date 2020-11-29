package taskconfig

import (
	"errors"
	"time"

	galaxyDB "github.com/galaxy-center/galaxy/lifecycle"
	models "github.com/galaxy-center/galaxy/models"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

// TaskConfig is an object representing the database table.
type TaskConfig struct {
	ID        uint64         `gorm:"primaryKey,autoIncrement" json:"id" toml:"id" yaml:"id"`
	Headers   datatypes.JSON `gorm:"type:json,column:headers" json:"headers" toml:"headers" yaml:"headers"`
	Content   datatypes.JSON `gorm:"type:json,column:content" json:"content" toml:"content" yaml:"content"`
	DeletedAt uint64         `gorm:"column:deleted_at" json:"deleted_at" toml:"deleted_at" yaml:"deleted_at"`
	CreatedAt uint64         `gorm:"autoCreateTime:nano" json:"created_at" toml:"created_at" yaml:"created_at"`
	CreatedBy string         `gorm:"column:created_by" json:"created_by,omitempty" toml:"created_by" yaml:"created_by,omitempty"`
	UpdatedAt uint64         `gorm:"autoUpdateTime:nano" json:"updated_at" toml:"updated_at" yaml:"updated_at"`
	UpdatedBy string         `gorm:"column:updated_by" json:"updated_by,omitempty" toml:"updated_by" yaml:"updated_by,omitempty"`
}

// TaskConfigColumns table field name.
var TaskConfigColumns = struct {
	ID        string
	Headers   string
	Content   string
	DeletedAt string
	CreatedAt string
	CreatedBy string
	UpdatedAt string
	UpdatedBy string
}{
	ID:        "id",
	Headers:   "headers",
	Content:   "content",
	DeletedAt: "deleted_at",
	CreatedAt: "created_at",
	CreatedBy: "created_by",
	UpdatedAt: "updated_at",
	UpdatedBy: "updated_by",
}

// AfterCreate do somethings, e.g. debug log.
func (t *TaskConfig) AfterCreate(tx *gorm.DB) (err error) {
	// nothing doing
	return
}

// Create a single TaskConfig to db by *gorm.DBs
func Create(config *TaskConfig) error {
	db := galaxyDB.GetDB()
	err := db.Create(config).Error
	return err
}

// BeforeUpdate do somethings, e.g. updating the updated_at value.
func (t *TaskConfig) BeforeUpdate(tx *gorm.DB) (err error) {
	t.UpdatedAt = uint64(time.Now().UnixNano())
	return
}

// AfterUpdate do somethings, e.g. update other database in the same transcation.
// 在同一个事务中更新数据
func (t *TaskConfig) AfterUpdate(tx *gorm.DB) (err error) {
	// nothing doing
	return
}

// Save a single task will be stored to db
// Note: all the fields will be updated to db, includes default value.
func Save(config *TaskConfig) error {
	db := galaxyDB.GetDB()
	err := db.Save(config).Error
	return err
}

// Updates updates from specific task that will not updating the zero value
// fields to db.
// 只能保存非零字段
func Updates(config *TaskConfig) error {
	db := galaxyDB.GetDB()
	err := db.Model(config).Updates(config).Error
	return err
}

// UpdatesFromMap updates from specific task that will not updating the zero value
// fields to db.
// 只能保存map包含字段
func UpdatesFromMap(id uint64, values map[string]interface{}) error {
	db := galaxyDB.GetDB()
	err := db.Model(&TaskConfig{}).Where("id = ?", id).Updates(values).Error
	return err
}

// Delete delete permanently. 永久删除
func Delete(id uint64) error {
	db := galaxyDB.GetDB()
	err := db.Delete(&TaskConfig{}, id).Error
	return err
}

// DeleteAt delete softly. 软删除
func DeleteAt(id uint64) error {
	db := galaxyDB.GetDB()
	err := db.Model(&TaskConfig{}).Where("id = ?", id).Update("deleted_at", time.Now().UnixNano()).Error
	return err
}

// Get returns the task by specific id.
func Get(id uint64) (*TaskConfig, error) {
	db := galaxyDB.GetDB()
	var config TaskConfig
	if err := db.First(&config, id).Error; err != nil {
		return nil, err
	}
	return &config, nil
}

// GetExcludeDeleted returns the task that excludes inactived by specific id.
func GetExcludeDeleted(id uint64) (*TaskConfig, error) {
	db := galaxyDB.GetDB()
	var config TaskConfig
	if err := db.Where("deleted_at = ?", 0).First(&config).Error; err != nil {
		return nil, err
	}
	return &config, nil
}

// JSONQuery returns TaskConfig array.
func JSONQuery(d models.InnerDetector) ([]TaskConfig, error) {
	db := galaxyDB.GetDB()

	var jqe *datatypes.JSONQueryExpression
	if d.GetLevel() == 1 {
		jqe = datatypes.JSONQuery(d.GetField()).Equals(d.GetValue(), d.GetKey1())
	} else if d.GetLevel() == 2 {
		jqe = datatypes.JSONQuery(d.GetField()).Equals(d.GetValue(), d.GetKey1(), d.GetKey2())
	} else {
		return nil, errors.New("Error InnerDetector")
	}

	var configs []TaskConfig
	db.Find(&configs, jqe)

	return configs, nil
}

// PaginateQuery todo
func PaginateQuery(p *models.Pagination) (models.Response, error) {
	var response models.Response
	response.Page = p.GetPage()

	db := galaxyDB.GetDB()

	var total int64
	attached := models.Attach(p.BuildCondition())

	db.Model(&TaskConfig{}).Scopes(attached).Count(&total)
	response.Total = int(total)
	response.TotalPage = int(total)/p.GetPageSize() + 1

	var configs []TaskConfig
	db.Scopes(attached, models.Paginate(p)).Find(&configs)
	response.Data = configs

	return response, nil
}
