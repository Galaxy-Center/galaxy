package models

import (
	"testing"

	db "github.com/galaxy-center/galaxy/lifecycle"
)

func TestCreate(t *testing.T) {
	db.Init()
	Create(&Task{
		ID:                 1,
		Name:               "test",
		Code:               "code",
		Type:               "JOB",
		Status:             "ENABLED",
		ExpiredAt:          100,
		Timeout:            3600,
		SchedulingType:     "RPC",
		SchedulingCategory: "RPC",
		Assess:             "Assess",
		Executor:           "Executor",
		Active:             true,
		CreatedAt:          10000,
		UpdatedAt:          10000,
	})
}
