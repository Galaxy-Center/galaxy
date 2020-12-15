package utils

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetIntOrDefault return parse int value, otherwise return def.
func GetIntOrDefault(c *gin.Context, key string, def int) int {
	v := c.Param(key)
	if v == "" {
		return def
	}
	pv, err := strconv.ParseInt(v, 10, 32)
	if err != nil {
		return def
	}
	return int(pv)
}

// GetIntOrPanic return parse int value, otherwise panic.
func GetIntOrPanic(c *gin.Context, key string, def int) int {
	v := c.Param(key)
	if v == "" {
		return def
	}
	pv, err := strconv.ParseInt(v, 10, 32)
	if err != nil {
		panic(fmt.Sprintf("can't parse %s to int value", v))
	}
	return int(pv)
}

// GetUint64OrDefault return parse int value, otherwise return def.
func GetUint64OrDefault(c *gin.Context, key string, def uint64) uint64 {
	v := c.Param(key)
	if v == "" {
		return def
	}
	pv, err := strconv.ParseUint(v, 10, 64)
	if err != nil {
		return def
	}
	return pv
}

// GetUint64OrPanic return parse int value, otherwise panic.
func GetUint64OrPanic(c *gin.Context, key string, def uint64) uint64 {
	v := c.Param(key)
	if v == "" {
		return def
	}
	pv, err := strconv.ParseUint(v, 10, 64)
	if err != nil {
		panic(fmt.Sprintf("can't parse %s to uint64 value", v))
	}
	return pv
}
