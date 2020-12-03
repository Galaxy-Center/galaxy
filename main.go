// Galaxy app support by Golang.
package main

import (
	"context"
	"net/http"
	"time"

	"github.com/galaxy-center/galaxy/config"
	logger "github.com/galaxy-center/galaxy/log"
	"github.com/galaxy-center/galaxy/resources"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	uuid "github.com/satori/go.uuid"
)

var (
	log     = logger.Get()
	rawLog  = logger.GetRaw()
	mainLog = log.WithField("prefix", "main")
)

func main() {
	Start()

	log.Info("galaxy server starting")

	router := gin.Default()
	registers(router)
	router.Run(":8080")
}

// Start do somethings about inializations.
func Start() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	config.SetNodeID("solo-" + uuid.NewV4().String())

	if err := initialiseSystem(ctx); err != nil {
		mainLog.Fatalf("Error initialising system: %v", err)
	}
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

func initialiseSystem(ctx context.Context) error {

	return nil
}
