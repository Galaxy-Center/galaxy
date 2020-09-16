package lifecycle

import (
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Database Wrapper, includes *grom.DB pools
type Database struct {
	*gorm.DB
}

var gormDB *gorm.DB

// Init initials a database and save the reference to `Database` struct.
func Init() {
	dsn := "lance:lancexu1992@tcp(45.32.253.249:3306)/galaxy?charset=utf8mb4&parseTime=true&loc=Local"
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       dsn,   // DSN data source name
		DefaultStringSize:         256,   // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据当前 MySQL 版本自动配置
	}), &gorm.Config{})
	if err != nil {
		panic("galaxy: occurred error connect to db")
	}
	gormDB = db

	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(10)           // 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxOpenConns(32)           // 设置打开数据库连接的最大数量
	sqlDB.SetConnMaxLifetime(time.Hour) // 设置连接可复用的最大时间
}

// GetDB returns *gorm.DB
func GetDB() *gorm.DB {
	if gormDB == nil {
		panic("galaxy: Database connection does not exist")
	}

	return gormDB
}
