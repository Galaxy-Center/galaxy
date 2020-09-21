package main

import (
	// "net/http"
	// "time"

	// "github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"

	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database/mysql"
	_ "github.com/golang-migrate/migrate/source/file"
)

func main() {
	log.Info("galaxy server starting")

	MigrateDB2()

	// router := gin.Default()
	// router.GET("/about", func(c *gin.Context) {
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"app":  "galaxy",
	// 		"time": time.Now().String(),
	// 	})
	// })
	// router.Run(":8080")
}

func MigrateDB2() {
	db, _ := sql.Open("mysql", "lance:Lancexu@1992@tcp(localhost:3306)/galaxy_test?charset=utf8mb4&parseTime=true&loc=Local&multiStatements=true")
	driver, _ := mysql.WithInstance(db, &mysql.Config{})
	m, _ := migrate.NewWithDatabaseInstance(
		"file:///migrations",
		"mysql",
		driver,
	)
	m.Steps(2)
}
