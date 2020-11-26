// Galaxy app support by Golang.
package main

import (
	"net/http"
	"time"

	logger "github.com/galaxy-center/galaxy/log"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

var (
	log    = logger.Get()
	rawLog = logger.GetRaw()
)

func main() {
	log.Info("galaxy server starting")

	router := gin.Default()
	router.GET("/about", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"app":  "galaxy",
			"time": time.Now().String(),
		})
	})
	router.Run(":8080")
}
