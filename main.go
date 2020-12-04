// Galaxy app support by Golang.
package main

import (
	"context"
	"net/http"
	"time"

	dbProvider "github.com/galaxy-center/galaxy/lifecycle"
	logger "github.com/galaxy-center/galaxy/log"
	"github.com/galaxy-center/galaxy/resources"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

const ()

var (
	log     = logger.Get()
	rawLog  = logger.GetRaw()
	mainLog = log.WithField("prefix", "main")
)

func main() {
	mainLog.Info("Galaxy Application starting.")
	router := gin.Default()
	registers(router)
	router.Run(":8080")
}

func registers(router *gin.Engine) {
	router.GET("/about", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"app":  "galaxy",
			"time": time.Now().String(),
		})
	})
	taskGroup := router.Group("/v1/task")
	taskGroup.POST("/", resources.CreateT)
	taskGroup.GET("/:id", resources.GetT)
}

func init() {
	_, cancel := context.WithCancel(context.Background())
	defer cancel()

	// if err := config.InitialiseSystem(); err != nil {
	// 	mainLog.Fatalf("Error initialising system: %v", err)
	// }

	dbProvider.Init()
}
