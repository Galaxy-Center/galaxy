package task

import (
	"os"
	"strconv"
	"testing"
	"time"

	db "github.com/galaxy-center/galaxy/lifecycle"
	migrateProvider "github.com/galaxy-center/galaxy/migrate"
	models "github.com/galaxy-center/galaxy/models"
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
		Type:               DelayQueue,
		Status:             ENABLED,
		ExpiredAt:          100,
		Timeout:            3600,
		SchedulingCategory: SINGLETON,
		Executor:           RPC,
		DeletedAt:          0,
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
		Type:               DelayJob,
		Status:             ENABLED,
		ExpiredAt:          100,
		Timeout:            3600,
		SchedulingCategory: SINGLETON,
		Executor:           RPC,
		DeletedAt:          0,
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
		Type:               DelayJob,
		Status:             ENABLED,
		ExpiredAt:          100,
		Timeout:            3600,
		SchedulingCategory: SINGLETON,
		Executor:           RPC,
		DeletedAt:          0,
		CreatedAt:          uint64(time.Now().UnixNano()),
		UpdatedAt:          uint64(time.Now().UnixNano()),
	}
	Create(task)

	task.ExpiredAt = 0
	task.Code = ""
	Updates(task)

	tmp, _ := Get(task.ID)
	assert.NotNil(t, tmp, "tmp should be not null")
	assert.EqualValues(t, 100, tmp.ExpiredAt, "expiredAt error")
	assert.EqualValues(t, "codeA", tmp.Code, "code error")
}

func TestUpdatesFromMap(t *testing.T) {
	m, _ := migrateProvider.BuildMigration()
	migrateProvider.Up(m)
	defer migrateProvider.Drop(m)

	task := &Task{
		Name:               "test",
		Code:               "codeA",
		Type:               DelayJob,
		Status:             ENABLED,
		ExpiredAt:          100,
		Timeout:            3600,
		SchedulingCategory: SINGLETON,
		Executor:           RPC,
		DeletedAt:          0,
		CreatedAt:          uint64(time.Now().UnixNano()),
		UpdatedAt:          uint64(time.Now().UnixNano()),
	}
	Create(task)

	UpdatesFromMap(task.ID, map[string]interface{}{"code": "codeB", "timeout": 0, "deleted_at": 100, "expired_at": 0})

	tmp, _ := Get(task.ID)
	assert.NotNil(t, tmp, "tmp should be not null")
	assert.EqualValues(t, 0, tmp.ExpiredAt, "expired_at error")
	assert.EqualValues(t, "codeB", tmp.Code, "code error")
	assert.EqualValues(t, 100, tmp.DeletedAt, "actived should is false")
}

func TestDelete(t *testing.T) {
	m, _ := migrateProvider.BuildMigration()
	migrateProvider.Up(m)
	defer migrateProvider.Drop(m)

	task := &Task{
		Name:               "test",
		Code:               "codeA",
		Type:               DelayJob,
		Status:             ENABLED,
		ExpiredAt:          100,
		Timeout:            3600,
		SchedulingCategory: SINGLETON,
		Executor:           RPC,
		DeletedAt:          0,
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

func TestDeleteAt(t *testing.T) {
	m, _ := migrateProvider.BuildMigration()
	migrateProvider.Up(m)
	defer migrateProvider.Drop(m)

	task := &Task{
		Name:               "test",
		Code:               "codeA",
		Type:               DelayJob,
		Status:             ENABLED,
		ExpiredAt:          100,
		Timeout:            3600,
		SchedulingCategory: SINGLETON,
		Executor:           RPC,
		DeletedAt:          0,
		CreatedAt:          uint64(time.Now().UnixNano()),
		UpdatedAt:          uint64(time.Now().UnixNano()),
	}
	Create(task)

	tmp, _ := GetExcludeDeleted(task.ID)
	assert.NotNil(t, tmp, "tmp should be not null")

	DeleteAt(task.ID)

	tmp2, _ := GetExcludeDeleted(task.ID)
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
			Type:               DelayJob,
			Status:             ENABLED,
			ExpiredAt:          uint64(100 * (i + 1)),
			Timeout:            3600 + (i * 100),
			SchedulingCategory: SINGLETON,
			Executor:           RPC,
			CreatedAt:          uint64(time.Now().UnixNano()),
			UpdatedAt:          uint64(time.Now().UnixNano()),
		}
		if i == 9 {
			task.DeletedAt = 900
		} else {
			task.DeletedAt = 0
		}
		if i%3 == 0 {
			task.Executor = RPC
		} else if i%3 == 1 {
			task.Executor = KAFKA
		} else {
			task.Executor = HTTP
		}
		Create(task)
	}

	p := new(models.Pagination)
	p.SetPage(1)
	p.SetPageSize(10)
	p.SetAttachment(models.Attachment{})
	p.GetAttachment()[models.PaginationColumns.Deleted] = true
	p.GetAttachment()[models.PaginationColumns.TimeRange] = models.Uint64Range{}.Default()
	p.GetAttachment()["code"] = "code5"

	res, _ := PaginateQuery(p)
	assert.NotNil(t, res, "res should not null")
	assert.EqualValues(t, res.Total, 1, "total error")
	assert.EqualValues(t, res.Data.([]Task)[0].ID, uint64(6), "ID error")
}
