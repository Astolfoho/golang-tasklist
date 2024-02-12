package task

import "task-list/internal/domain/taskconfiguration"

type Task struct {
	taskconfiguration.TaskConfiguration
	Finished bool
}
