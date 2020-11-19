package models

import (
	// "os"
	"testing"
	"time"

	db "github.com/galaxy-center/galaxy/lifecycle"
	migrate_provider "github.com/galaxy-center/galaxy/migrate"
	"github.com/golang-migrate/migrate/v4"
	log "github.com/sirupsen/logrus"
)

// func TestCreateDB(t *testing.T) {
// 	db.Init()

// 	db.GetDB().AutoMigrate(&Task{})
// }

func TestCreate(t *testing.T) {
	m, _ := migrate_provider.BuildMigration()

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatalf("An error occurred while syncing the database... %v", err)
	}

	log.Println("Database migrated")

	// os.Exit(0)
	// migrate.ExecuteDrop()

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

	m, _ = migrate_provider.BuildMigration()
	if err := m.Drop(); err != nil && err != migrate.ErrNoChange {
		log.Fatalf("An error occurred while syncing the database... %v", err)
	}

	log.Println("Database droped")

	// os.Exit(0)
}
