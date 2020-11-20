package models

import (
	"os"
	"strconv"
	"testing"
	"time"

	db "github.com/galaxy-center/galaxy/lifecycle"
	migrateProvider "github.com/galaxy-center/galaxy/migrate"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	db.Init()
	code := m.Run()
	os.Exit(code)
}

func TestCreate(t *testing.T) {
	// Integrated database structure migration.
	m, _ := migrateProvider.BuildMigration()
	migrateProvider.Up(m)
	defer migrateProvider.Drop(m)

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
		CreatedAt:          uint64(time.Now().UnixNano()),
		UpdatedAt:          uint64(time.Now().UnixNano()),
	}
	Create(task)

	exist, _ := Get(1)
	assert.NotNil(t, exist, "exist should be not null")
	assert.EqualValues(t, 100, exist.ExpiredAt, "err")
}

func TestSave(t *testing.T) {
	m, _ := migrateProvider.BuildMigration()
	migrateProvider.Up(m)
	defer migrateProvider.Drop(m)

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
		CreatedAt:          uint64(time.Now().UnixNano()),
		UpdatedAt:          uint64(time.Now().UnixNano()),
	}
	Create(task)

	task.ExpiredAt = 0
	task.Code = "codeB"
	Save(task)

	tmp, _ := Get(task.ID)
	assert.NotNil(t, tmp, "tmp should be not null")
	assert.EqualValues(t, 0, tmp.ExpiredAt)
	assert.EqualValues(t, "codeB", tmp.Code)
}

func TestUpdates(t *testing.T) {
	m, _ := migrateProvider.BuildMigration()
	migrateProvider.Up(m)
	defer migrateProvider.Drop(m)

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
		CreatedAt:          uint64(time.Now().UnixNano()),
		UpdatedAt:          uint64(time.Now().UnixNano()),
	}
	Create(task)

	task.ExpiredAt = 0
	task.Code = ""
	Updates(task)

	tmp, _ := Get(task.ID)
	assert.NotNil(t, tmp, "tmp should be not null")
	assert.EqualValues(t, 100, tmp.ExpiredAt)
	assert.EqualValues(t, "codeA", tmp.Code)
}

func TestUpdatesFromMap(t *testing.T) {
	m, _ := migrateProvider.BuildMigration()
	migrateProvider.Up(m)
	defer migrateProvider.Drop(m)

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
		CreatedAt:          uint64(time.Now().UnixNano()),
		UpdatedAt:          uint64(time.Now().UnixNano()),
	}
	Create(task)

	UpdatesFromMap(task.ID, map[string]interface{}{"code": "codeB", "timeout": 0, "actived": false, "expired_at": 0})

	tmp, _ := Get(task.ID)
	assert.NotNil(t, tmp, "tmp should be not null")
	assert.EqualValues(t, 0, tmp.ExpiredAt, "expired_at error")
	assert.EqualValues(t, "codeB", tmp.Code, "code error")
	assert.False(t, tmp.Actived, "actived should is false")
}

func TestDelete(t *testing.T) {
	m, _ := migrateProvider.BuildMigration()
	migrateProvider.Up(m)
	defer migrateProvider.Drop(m)

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
		CreatedAt:          uint64(time.Now().UnixNano()),
		UpdatedAt:          uint64(time.Now().UnixNano()),
	}
	Create(task)

	tmp, _ := Get(task.ID)
	assert.NotNil(t, tmp, "tmp should be not null")

	Delete(task.ID)

	tmp2, _ := Get(task.ID)
	assert.Nil(t, tmp2, "tmp shoule be null after deleted")
}

func TestDeleteByActived(t *testing.T) {
	m, _ := migrateProvider.BuildMigration()
	migrateProvider.Up(m)
	defer migrateProvider.Drop(m)

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
		CreatedAt:          uint64(time.Now().UnixNano()),
		UpdatedAt:          uint64(time.Now().UnixNano()),
	}
	Create(task)

	tmp, _ := GetExcludeInactived(task.ID)
	assert.NotNil(t, tmp, "tmp should be not null")

	DeleteByActived(task.ID)

	tmp2, _ := GetExcludeInactived(task.ID)
	assert.Nil(t, tmp2, "tmp should be null after deleted")
}

func TestPaginateQuery(t *testing.T) {
	m, _ := migrateProvider.BuildMigration()
	migrateProvider.Up(m)
	defer migrateProvider.Drop(m)

	for i := 0; i < 10; i++ {
		task := &Task{
			Name:               "test" + strconv.Itoa(i),
			Code:               "code" + strconv.Itoa(i),
			Type:               "JOB",
			Status:             "ENABLED",
			ExpiredAt:          uint64(100 * (i + 1)),
			Timeout:            3600 + (i * 100),
			SchedulingType:     "RPC",
			SchedulingCategory: "RPC",
			Assess:             "Assess",
			Executor:           "Executor",
			Actived:            true,
			CreatedAt:          uint64(time.Now().UnixNano()),
			UpdatedAt:          uint64(time.Now().UnixNano()),
		}
		Create(task)
	}
}
