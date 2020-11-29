// Package migrate support by golang-migrate.
package migrate

import (
	"fmt"

	log "github.com/sirupsen/logrus"

	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

// BuildMigration returns *migrate.Migrate if everything is ok. otherwise return specified error.
func BuildMigration() (*migrate.Migrate, error) {
	db, err := sql.Open(
		"mysql",
		"lance:Lancexu@1992@tcp(localhost:3306)/galaxy_test?charset=utf8mb4&parseTime=true&loc=Local&multiStatements=true")
	if err != nil {
		log.Fatalf("couldn't connect to the Mysql database... %v", err)
	}
	if err := db.Ping(); err != nil {
		log.Fatalf("couldn't not ping DB... %v", err)
	}

	driver, err := mysql.WithInstance(db, &mysql.Config{})
	if err != nil {
		log.Fatalf("couldn't not start sql migration... %v", err)
	}
	m, err := migrate.NewWithDatabaseInstance(
		fmt.Sprintf("file://%s", "../migrationfiles"),
		"mysql",
		driver,
	)
	if err != nil {
		log.Fatalf("migration failed... %v", err)
		return nil, err
	}
	return m, nil
}

// Up execute up migrations.
func Up(m *migrate.Migrate) {
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatalf("An error occurred while syncing the database... %v", err)
	}

	log.Println("Database migrated")
}

// Drop deletes everything in the database.
func Drop(m *migrate.Migrate) {
	if err := m.Drop(); err != nil && err != migrate.ErrNoChange {
		log.Fatalf("An error occurred while syncing the database... %v", err)
	}

	log.Println("Database dropped")

	close(m)
}

// close close the *migrate.Migrate.
func close(m *migrate.Migrate) {
	sourceErr, databaseErr := m.Close()
	if sourceErr != nil {
		log.Fatalf("An source error occurred while close migrate... %v", sourceErr)
	}
	if databaseErr != nil {
		log.Fatalf("An database error occurred while close migrate... %v", databaseErr)
	}

	log.Println("Migrate closed")
}
