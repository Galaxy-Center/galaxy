package config

import (
	"fmt"
)

const (
	// MySQLDSNFormat string format
	MySQLDSNFormat = "%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=true&loc=Local"
)

// DBConfig defines for DB connection.
type DBConfig struct {
	User     string
	Password string
	Host     string
	Port     int
	Database string
}

func (c *DBConfig) format() string {
	return fmt.Sprintf(MySQLDSNFormat, c.User, c.Password, c.Host, c.Port, c.Database)
}
