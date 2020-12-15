package services

import (
	"fmt"
	"net/http"

	"github.com/galaxy-center/galaxy/commons"
	"github.com/galaxy-center/galaxy/models"
	"github.com/galaxy-center/galaxy/models/task"
)

// CreateTask returns status.
func CreateTask(t *task.Task) *commons.Error {
	if err := task.Create(t); err != nil {
		log.WithField("task", t).Errorf("occurred exception when inserting task: %v", err)
		return commons.StatusDBOperationAbnormal
	}
	return nil
}

// UpdateTask returns error.
func UpdateTask(t *task.Task) *commons.Error {
	if err := task.Updates(t); err != nil {
		log.WithField("task", t).Errorf("occurred exception when updating task: %v", err)
		return commons.StatusDBOperationAbnormal
	}
	return nil
}

// UpsertTask create or update.
func UpsertTask(t *task.Task) *commons.Error {
	if t.ID == uint64(0) {
		return CreateTask(t)
	}
	return UpdateTask(t)
}

// DeleteTask delete by os storage.
func DeleteTask(id uint64, deletedAt bool) (bool, *commons.Error) {
	if deletedAt {
		if err := task.DeleteAt(id); err != nil {
			log.Errorf("occurred exception when deleting task: %d", id)
			return false, commons.StatusDBOperationAbnormal
		}
		return true, nil
	}
	if err := task.Delete(id); err != nil {
		log.Errorf("occurred exception when deleting task: %d", id)
		return false, commons.StatusDBOperationAbnormal
	}
	return true, nil
}

// GetTask returns target or status.
func GetTask(id uint64) (*task.Task, *commons.Error) {
	t, err := task.Get(id)
	if err != nil {
		log.WithField("id", id).Error("occurred exception when getting task")
		return nil, commons.StatusDBOperationAbnormal
	}
	if t == nil {
		return nil, &commons.Error{
			Code:  http.StatusNotFound,
			Error: fmt.Errorf("Not found %d", id)}
	}
	return t, nil
}

// GetTasksWith pagination queries.
func GetTasksWith(p *models.Pagination) (*models.Response, *commons.Error) {
	res, err := task.PaginateQuery(p)
	if err != nil {
		log.WithField("pagination", p).Error("occurred exception when getting tasks")
		return nil, commons.StatusDBOperationAbnormal
	}

	return &res, nil
}
