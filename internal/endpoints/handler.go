package endpoints

import "task-list/internal/domain/taskconfiguration"

type Handler struct {
	TaskConfigurationService taskconfiguration.Service
}
