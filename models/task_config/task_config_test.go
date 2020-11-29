package taskconfig

import (
	"encoding/json"
	"os"
	"testing"
	"time"

	db "github.com/galaxy-center/galaxy/lifecycle"
	migrateProvider "github.com/galaxy-center/galaxy/migrate"
	models "github.com/galaxy-center/galaxy/models"
	"github.com/stretchr/testify/assert"
	"gorm.io/datatypes"
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

	foo := make(map[string]interface{})
	foo["foo"] = "bar"
	config := &TaskConfig{
		Headers:   datatypes.JSON([]byte(`{"name": "jinzhu", "age": 18, "tags": ["tag1", "tag2"], "orgs": {"orga": "orga"}}`)),
		Content:   datatypes.JSON([]byte(`{"name": "jinzhu", "age": 18, "tags": ["tag1", "tag2"], "orgs": {"orga": "orga"}}`)),
		DeletedAt: 0,
		CreatedAt: uint64(time.Now().UnixNano()),
		UpdatedAt: uint64(time.Now().UnixNano()),
	}
	Create(config)

	exist, _ := Get(1)
	assert.NotNil(t, exist, "exist should be not null")
	assert.EqualValues(t, 0, exist.DeletedAt, "err")
	assert.True(t, exist.CreatedAt < uint64(time.Now().UnixNano()), "time error")
}

func TestSave(t *testing.T) {
	m, _ := migrateProvider.BuildMigration()
	migrateProvider.Up(m)
	defer migrateProvider.Drop(m)

	config := &TaskConfig{
		Headers:   datatypes.JSON([]byte(`{"name": "jinzhu", "age": 18, "tags": ["tag1", "tag2"], "orgs": {"orga": "orga"}}`)),
		Content:   datatypes.JSON([]byte(`{"name": "jinzhu", "age": 18, "tags": ["tag1", "tag2"], "orgs": {"orga": "orga"}}`)),
		DeletedAt: 0,
		CreatedAt: uint64(time.Now().UnixNano()),
		UpdatedAt: uint64(time.Now().UnixNano()),
	}
	Create(config)

	var a map[string]interface{}
	json.Unmarshal(config.Headers, &a)
	a["name"] = "lance"
	b, _ := json.Marshal(a)
	config.Headers = datatypes.JSON(b)
	Save(config)

	tmp, _ := Get(config.ID)
	assert.NotNil(t, tmp, "tmp should be not null")
	json.Unmarshal(config.Headers, &a)
	assert.EqualValues(t, "lance", a["name"], "json field update failed")
}

func TestUpdates(t *testing.T) {
	m, _ := migrateProvider.BuildMigration()
	migrateProvider.Up(m)
	defer migrateProvider.Drop(m)

	config := &TaskConfig{
		Headers:   datatypes.JSON([]byte(`{"name": "jinzhu", "age": 18, "tags": ["tag1", "tag2"], "orgs": {"orga": "orga"}}`)),
		Content:   datatypes.JSON([]byte(`{"name": "jinzhu", "age": 18, "tags": ["tag1", "tag2"], "orgs": {"orga": "orga"}}`)),
		DeletedAt: 0,
		CreatedAt: uint64(time.Now().UnixNano()),
		UpdatedAt: uint64(time.Now().UnixNano()),
	}
	Create(config)

	config.Headers = nil
	config.Content = nil
	Updates(config)

	tmp, _ := Get(config.ID)
	assert.NotNil(t, tmp, "tmp should be not null")
	assert.NotNil(t, tmp.Headers, "headers error")
	assert.NotNil(t, tmp.Content, "content error")
}

func TestUpdatesFromMap(t *testing.T) {
	m, _ := migrateProvider.BuildMigration()
	migrateProvider.Up(m)
	defer migrateProvider.Drop(m)

	config := &TaskConfig{
		Headers:   datatypes.JSON([]byte(`{"name": "jinzhu", "age": 18, "tags": ["tag1", "tag2"], "orgs": {"orga": "orga"}}`)),
		Content:   datatypes.JSON([]byte(`{"name": "jinzhu", "age": 18, "tags": ["tag1", "tag2"], "orgs": {"orga": "orga"}}`)),
		DeletedAt: 0,
		CreatedAt: uint64(time.Now().UnixNano()),
		UpdatedAt: uint64(time.Now().UnixNano()),
	}
	Create(config)

	b, _ := json.Marshal(map[string]interface{}{"modify": map[string]string{"foo": "bar"}})
	UpdatesFromMap(config.ID, map[string]interface{}{
		"headers":    datatypes.JSON(b),
		"deleted_at": uint64(100),
	})

	tmp, _ := Get(config.ID)
	assert.NotNil(t, tmp, "tmp should be not null")
	assert.EqualValues(t, uint64(100), tmp.DeletedAt, "actived should is false")
}

func TestDelete(t *testing.T) {
	m, _ := migrateProvider.BuildMigration()
	migrateProvider.Up(m)
	defer migrateProvider.Drop(m)

	config := &TaskConfig{
		Headers:   datatypes.JSON([]byte(`{"name": "jinzhu", "age": 18, "tags": ["tag1", "tag2"], "orgs": {"orga": "orga"}}`)),
		Content:   datatypes.JSON([]byte(`{"name": "jinzhu", "age": 18, "tags": ["tag1", "tag2"], "orgs": {"orga": "orga"}}`)),
		DeletedAt: 0,
		CreatedAt: uint64(time.Now().UnixNano()),
		UpdatedAt: uint64(time.Now().UnixNano()),
	}
	Create(config)

	tmp, _ := Get(config.ID)
	assert.NotNil(t, tmp, "tmp should be not null")

	Delete(config.ID)

	tmp2, _ := Get(config.ID)
	assert.Nil(t, tmp2, "tmp shoule be null after deleted")
}

func TestDeleteAt(t *testing.T) {
	m, _ := migrateProvider.BuildMigration()
	migrateProvider.Up(m)
	defer migrateProvider.Drop(m)

	config := &TaskConfig{
		Headers:   datatypes.JSON([]byte(`{"name": "jinzhu", "age": 18, "tags": ["tag1", "tag2"], "orgs": {"orga": "orga"}}`)),
		Content:   datatypes.JSON([]byte(`{"name": "jinzhu", "age": 18, "tags": ["tag1", "tag2"], "orgs": {"orga": "orga"}}`)),
		DeletedAt: 0,
		CreatedAt: uint64(time.Now().UnixNano()),
		UpdatedAt: uint64(time.Now().UnixNano()),
	}
	Create(config)

	tmp, _ := GetExcludeDeleted(config.ID)
	assert.NotNil(t, tmp, "tmp should be not null")

	DeleteAt(config.ID)

	tmp2, _ := GetExcludeDeleted(config.ID)
	assert.Nil(t, tmp2, "tmp should be null after deleted")
}

func TestPaginateQuery(t *testing.T) {
	m, _ := migrateProvider.BuildMigration()
	migrateProvider.Up(m)
	defer migrateProvider.Drop(m)

	for i := 0; i < 15; i++ {
		h := make(map[string]interface{})
		h["name"] = "lance"

		info := make(map[string]int)
		info["age"] = i

		h["info"] = info

		b, _ := json.Marshal(h)
		config := &TaskConfig{
			Headers:   datatypes.JSON([]byte(b)),
			Content:   datatypes.JSON([]byte(`{"name": "jinzhu", "age": 18, "tags": ["tag1", "tag2"], "orgs": {"orga": "orga"}}`)),
			DeletedAt: 0,
			CreatedAt: uint64(time.Now().UnixNano()),
			UpdatedAt: uint64(time.Now().UnixNano()),
		}
		if i == 9 {
			config.DeletedAt = 900
		} else {
			config.DeletedAt = 0
		}

		Create(config)
	}

	p := new(models.Pagination)
	p.SetPage(1)
	p.SetPageSize(10)
	p.SetAttachment(models.Attachment{})
	p.GetAttachment()[models.PaginationColumns.Deleted] = true
	p.GetAttachment()[models.PaginationColumns.TimeRange] = models.Uint64Range{}.Set(uint64(0), uint64(time.Now().UnixNano()))

	res, _ := PaginateQuery(p)
	assert.NotNil(t, res, "res should not null")
	assert.EqualValues(t, 14, res.Total, "total error")
	assert.EqualValues(t, 2, res.TotalPage, "totalPage error")
}

func TestJsonQuery(t *testing.T) {
	m, _ := migrateProvider.BuildMigration()
	migrateProvider.Up(m)
	defer migrateProvider.Drop(m)

	for i := 0; i < 15; i++ {
		h := make(map[string]interface{})
		h["name"] = "lance"

		info := make(map[string]int)
		info["age"] = i

		h["info"] = info

		b, _ := json.Marshal(h)
		config := &TaskConfig{
			Headers:   datatypes.JSON([]byte(b)),
			Content:   datatypes.JSON([]byte(`{"name": "jinzhu", "age": 18, "tags": ["tag1", "tag2"], "orgs": {"orga": "orga"}}`)),
			DeletedAt: 0,
			CreatedAt: uint64(time.Now().UnixNano()),
			UpdatedAt: uint64(time.Now().UnixNano()),
		}
		if i == 9 {
			config.DeletedAt = 900
		} else {
			config.DeletedAt = 0
		}

		Create(config)
	}

	configs, err := JSONQuery(models.InnerDetector{}.SetLevel2("headers", "info", "age", 14))
	assert.Nil(t, err)
	assert.EqualValues(t, 15, configs[0].ID, "error ID")
}
