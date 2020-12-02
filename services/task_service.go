package services

import (
	"github.com/galaxy-center/galaxy/commons"
	"github.com/galaxy-center/galaxy/models/task"
)

// CreateTask returns status.
func CreateTask(t *task.Task) commons.Status {
	if err := task.Create(t); err != nil {
		log.WithField("task", t).Error("occurred exception when inserting task", err)
		return commons.StatusDBOperationAbnormal
	}
	return commons.StatusOK
}

// GetTask returns target or status.
func GetTask(id uint64) (task.Task, commons.Status) {
	t, err := task.Get(id)
	if err != nil {
		log.WithField("id", id).Error("occurred exception when getting task")
		return task.Task{}, commons.StatusDBOperationAbnormal
	}
	if t == nil {
		return task.Task{}, commons.StatusNotFound
	}
	return *t, commons.StatusOK
}
