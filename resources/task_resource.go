package resources

import (
	"github.com/galaxy-center/galaxy/models/task"
)

// CreateT create func.
func CreateT(t *task.Task) WebAPIResponse {
	err := task.Create(t)
	if err != nil {
		return Success(t)
	}
	return Error(NOTFOUND, "400")
}
