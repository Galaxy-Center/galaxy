package resources

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/galaxy-center/galaxy/commons"
	task "github.com/galaxy-center/galaxy/models/task"
	services "github.com/galaxy-center/galaxy/services"
	"github.com/gin-gonic/gin"
)

// CreateT create func.
func CreateT(c *gin.Context) {
	var t task.Task
	c.BindJSON(&t)
	s := services.CreateTask(&t)
	if s.Code != commons.OK {
		c.JSON(http.StatusBadRequest, commons.Error(s))
		return
	}
	c.JSON(http.StatusOK, commons.Success(t))
}

// GetT get task by id.
func GetT(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, commons.Error(commons.StatusBadRequest))
		return
	}
	tid, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, commons.ErrorWithMessage(commons.BadRequest, fmt.Sprintf("%s invalid.", id)))
		return
	}
	exist, s := services.GetTask(tid)
	if s.Code != commons.OK {
		c.JSON(http.StatusFound, commons.Error(s))
		return
	}
	c.JSON(http.StatusOK, commons.Success(exist))
}
