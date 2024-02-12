package database

import (
	"task-list/internal/domain/taskconfiguration"

	"gorm.io/gorm"
)

type TaskConfigurationRepository struct {
	Db *gorm.DB
}

func (r *TaskConfigurationRepository) Create(t *taskconfiguration.TaskConfiguration) error {
	tx := r.Db.Create(t)
	return tx.Error
}

func (r *TaskConfigurationRepository) GetAll() ([]taskconfiguration.TaskConfiguration, error) {
	var ret []taskconfiguration.TaskConfiguration
	tx := r.Db.Find(&ret)
	return ret, tx.Error
}

func (r *TaskConfigurationRepository) GetById(id string) (*taskconfiguration.TaskConfiguration, error) {
	var ret taskconfiguration.TaskConfiguration
	tx := r.Db.First(&ret, "id = ?", id)
	return &ret, tx.Error
}
