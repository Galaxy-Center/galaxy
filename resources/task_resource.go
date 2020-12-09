//Package resources saves all APIs files.
// will filter out invalid request paramas checks work. status is 4xx.
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

	if err := services.CreateTask(&t); err != nil {
		c.JSON(err.Code, commons.ErrorWithMessage(err.Format()))
		return
	}
	log.WithField("task", t).Info("inserted a task")
	c.JSON(http.StatusOK, commons.Success(t))
}

// UpdateT update func.
func UpdateT(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, commons.ErrorWithMessage("param id invalid"))
		return
	}
	tid, err := strconv.ParseUint(id, 10, 64)
	if err != nil || tid <= 0 {
		c.JSON(
			http.StatusBadRequest,
			commons.ErrorWithMessage(fmt.Sprintf("%s invalid.", id)))
		return
	}

	var t task.Task
	c.BindJSON(&t)
	t.ID = tid

	if ce := services.UpsertTask(&t); ce != nil {
		c.JSON(ce.Code, commons.ErrorWithMessage(ce.Format()))
		return
	}
	c.JSON(http.StatusOK, commons.Success(t))
}

// DeleteT deleted options.
func DeleteT(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, commons.ErrorWithMessage("params invalid"))
		return
	}
	tid, err := strconv.ParseUint(id, 10, 64)
	if err != nil || tid <= 0 {
		c.JSON(
			http.StatusBadRequest,
			commons.ErrorWithMessage(fmt.Sprintf("%s invalid.", id)))
		return
	}

	if _, ce := services.DeleteTask(tid, true); ce != nil {
		c.JSON(
			ce.Code,
			commons.ErrorWithMessage(ce.Format()))
		return
	}
	c.JSON(http.StatusOK, commons.Success(tid))
}

// GetT get task by id.
func GetT(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, commons.ErrorWithMessage("params invalid"))
		return
	}
	tid, err := strconv.ParseUint(id, 10, 64)
	if err != nil || tid <= 0 {
		c.JSON(
			http.StatusBadRequest,
			commons.ErrorWithMessage(fmt.Sprintf("%s invalid.", id)))
		return
	}

	t, ce := services.GetTask(tid)
	if ce != nil {
		c.JSON(ce.Code, commons.ErrorWithMessage(ce.Format()))
		return
	}
	c.JSON(http.StatusOK, commons.Success(t))
}

// GetTWith query pagination.
func GetTWith(c *gin.Context) {
	// c.Param("")
	// task.PaginateQuery()
}
