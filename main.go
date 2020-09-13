package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
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
