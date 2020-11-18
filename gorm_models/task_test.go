package models

import (
	"testing"
	"time"

	db "github.com/galaxy-center/galaxy/lifecycle"
)

// func TestCreateDB(t *testing.T) {
// 	db.Init()

// 	db.GetDB().AutoMigrate(&Task{})
// }

func TestCreate(t *testing.T) {
	db.Init()

	task := &Task{
		Name:               "test",
		Code:               "codeA",
		Type:               "JOB",
		Status:             "ENABLED",
		ExpiredAt:          100,
		Timeout:            3600,
		SchedulingType:     "RPC",
		SchedulingCategory: "RPC",
		Assess:             "Assess",
		Executor:           "Executor",
		Actived:            true,
		CreatedAt:          10000,
		UpdatedAt:          10000,
	}
	Create(task)

	task.ExpiredAt = uint64(time.Now().UnixNano())
	Save(task)
}
