// Package migrate support by golang-migrate.
package migrate

import (
	"flag"
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"

	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

// buildMigration returns *migrate.Migrate if everything is ok. otherwise return specified error.
func buildMigration() (*migrate.Migrate, error) {
	var migrationDir = flag.String("migration.files", "./migrations", "Directory where the migration files are located ?")
	flag.Parse()

	db, err := sql.Open("mysql", "lance:Lancexu@1992@tcp(localhost:3306)/galaxy_test?charset=utf8mb4&parseTime=true&loc=Local&multiStatements=true")
	if err != nil {
		log.Fatalf("countn't connect to the Mysql database... %v", err)
	}
	if err := db.Ping(); err != nil {
		log.Fatalf("countn't not ping DB... %v", err)
	}

	driver, err := mysql.WithInstance(db, &mysql.Config{})
	if err != nil {
		log.Fatalf("countn't not start sql migration... %v", err)
	}
	m, err := migrate.NewWithDatabaseInstance(
		fmt.Sprintf("file://%s", *migrationDir),
		"mysql",
		driver,
	)
	if err != nil {
		log.Fatalf("migration falied... %v", err)
		return nil, err
	}
	return m, nil
}

// ExecuteUp execute up migrations.
func ExecuteUp() {
	m, err := buildMigration()
	if err != nil {
		return
	}
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatalf("An error occurred while syncing the database... %v", err)
	}

	log.Println("Database migrated")

	os.Exit(0)
}

// ExecuteDrop deletes everything in the database.
func ExecuteDrop() {
	m, err := buildMigration()
	if err != nil {
		return
	}
	if err := m.Drop(); err != nil && err != migrate.ErrNoChange {
		log.Fatalf("An error occurred while syncing the database... %v", err)
	}

	log.Println("Database droped")

	os.Exit(0)
}
