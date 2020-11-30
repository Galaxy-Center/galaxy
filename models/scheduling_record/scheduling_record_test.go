package schedulingrecord

import (
	"os"
	"testing"

	db "github.com/galaxy-center/galaxy/lifecycle"
	log "github.com/galaxy-center/galaxy/log"
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
	defer func() {
		err := recover()
		if err != nil {
			log.Get().Error("Occurred error:", err)
		}
		migrateProvider.Drop(m)
	}()

	record := &SchedulingRecord{
		TaskID:    uint64(1),
		Status:    NEW,
		Message:   "successful",
		DeletedAt: 0,
	}
	Create(record)

	exist, _ := Get(1)
	assert.NotNil(t, exist, "exist should be not null")
	assert.EqualValues(t, NEW, exist.Status, "status err")
}

func TestSave(t *testing.T) {
	m, _ := migrateProvider.BuildMigration()
	migrateProvider.Up(m)
	defer func() {
		err := recover()
		if err != nil {
			log.Get().Error("Occurred error:", err)
		}
		migrateProvider.Drop(m)
	}()

	record := &SchedulingRecord{
		TaskID:    uint64(1),
		Status:    NEW,
		Message:   "successful",
		DeletedAt: 0,
	}
	Create(record)

	record.Status = RUNNABLE
	record.Message = ""
	Save(record)

	tmp, _ := Get(record.ID)
	assert.NotNil(t, tmp, "tmp should be not null")
	assert.EqualValues(t, "", record.Message, "message error")
	assert.EqualValues(t, RUNNABLE, tmp.Status)
}

func TestUpdates(t *testing.T) {
	m, _ := migrateProvider.BuildMigration()
	migrateProvider.Up(m)
	defer func() {
		err := recover()
		if err != nil {
			log.Get().Error("Occurred error:", err)
		}
		migrateProvider.Drop(m)
	}()

	record := &SchedulingRecord{
		TaskID:    uint64(1),
		Status:    NEW,
		Message:   "successful",
		DeletedAt: 0,
	}
	Create(record)

	record.Status = RUNNABLE
	record.Message = "failed"
	Updates(record)

	tmp, _ := Get(record.ID)
	assert.NotNil(t, tmp, "tmp should be not null")
	assert.EqualValues(t, "failed", record.Message, "message error")
	assert.EqualValues(t, RUNNABLE, tmp.Status)
}

func TestUpdatesFromMap(t *testing.T) {
	m, _ := migrateProvider.BuildMigration()
	migrateProvider.Up(m)
	defer func() {
		err := recover()
		if err != nil {
			log.Get().Error("Occurred error:", err)
		}
		migrateProvider.Drop(m)
	}()

	record := &SchedulingRecord{
		TaskID:    uint64(1),
		Status:    NEW,
		Message:   "successful",
		DeletedAt: 0,
	}
	Create(record)

	UpdatesFromMap(record.ID, map[string]interface{}{"status": RUNNING, "task_id": 2})

	tmp, _ := Get(record.ID)
	assert.NotNil(t, tmp, "tmp should be not null")
	assert.EqualValues(t, RUNNING, tmp.Status, "status error")
	assert.EqualValues(t, 2, tmp.TaskID, "task_id error")
}

func TestDelete(t *testing.T) {
	m, _ := migrateProvider.BuildMigration()
	migrateProvider.Up(m)
	defer func() {
		err := recover()
		if err != nil {
			log.Get().Error("Occurred error:", err)
		}
		migrateProvider.Drop(m)
	}()

	record := &SchedulingRecord{
		TaskID:    uint64(1),
		Status:    NEW,
		Message:   "successful",
		DeletedAt: 0,
	}
	Create(record)

	tmp, _ := Get(record.ID)
	assert.NotNil(t, tmp, "tmp should be not null")

	Delete(record.ID)

	tmp2, _ := Get(record.ID)
	assert.Nil(t, tmp2, "tmp shoule be null after deleted")
}

func TestDeleteAt(t *testing.T) {
	m, _ := migrateProvider.BuildMigration()
	migrateProvider.Up(m)
	defer func() {
		err := recover()
		if err != nil {
			log.Get().Error("Occurred error:", err)
		}
		migrateProvider.Drop(m)
	}()

	record := &SchedulingRecord{
		TaskID:    uint64(1),
		Status:    NEW,
		Message:   "successful",
		DeletedAt: 0,
	}
	Create(record)

	tmp, _ := GetExcludeDeleted(record.ID)
	assert.NotNil(t, tmp, "tmp should be not null")

	DeleteAt(record.ID)

	tmp2, _ := GetExcludeDeleted(record.ID)
	assert.Nil(t, tmp2, "tmp should be null after deleted")
}

func TestPaginateQuery(t *testing.T) {
	m, _ := migrateProvider.BuildMigration()
	migrateProvider.Up(m)
	defer func() {
		err := recover()
		if err != nil {
			log.Get().Error("Occurred error:", err)
		}
		migrateProvider.Drop(m)
	}()

	for i := 0; i < 10; i++ {
		record := &SchedulingRecord{
			TaskID:    uint64(1),
			Status:    NEW,
			Message:   "successful",
			DeletedAt: 0,
		}
		if i == 0 {
			record.DeletedAt = 900
		} else {
			record.DeletedAt = 0
		}
		if i%3 == 0 {
			record.Status = NEW
		} else if i%3 == 1 {
			record.Status = RUNNABLE
		} else {
			record.Status = RUNNING
		}
		Create(record)
	}

	p := new(models.Pagination)
	p.SetPage(1)
	p.SetPageSize(10)
	p.SetAttachment(models.Attachment{})
	p.GetAttachment()[models.PaginationColumns.Deleted] = true
	p.GetAttachment()[models.PaginationColumns.TimeRange] = models.Uint64Range{}.Default()
	p.GetAttachment()["status"] = RUNNABLE

	res, _ := PaginateQuery(p)
	assert.NotNil(t, res, "res should not null")
	assert.EqualValues(t, 3, res.Total, "total error")
	assert.EqualValues(t, 3, len(res.Data.([]SchedulingRecord)), "data size error")
	assert.EqualValues(t, res.Data.([]SchedulingRecord)[0].ID, uint64(2), "ID error")
}
