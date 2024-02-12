package taskconfiguration

import (
	internalerrors "task-list/internal/domain/internal-errors"
	"task-list/internal/domain/taskconfiguration/contract"
)

type Service interface {
	Create(tc *contract.NewTaskConfiguration) (id string, err error)
	GetAll() ([]TaskConfiguration, error)
	GetById(id string) (taskConfiguration *TaskConfiguration, err error)
}

type ServiceImp struct {
	Repository Repository
}

func (s ServiceImp) Create(tc *contract.NewTaskConfiguration) (id string, err error) {

	taskConfig, err := NewTaskConfig(tc.Name, tc.Description, Weekdays(tc.DaysOfWeek))

	if err != nil {
		return "", err
	}

	err = s.Repository.Create(taskConfig)

	if err != nil {
		return "", internalerrors.ErrInternalServerError
	}

	return taskConfig.Id, nil
}

func (s ServiceImp) GetAll() ([]TaskConfiguration, error) {
	return s.Repository.GetAll()
}

func (s ServiceImp) GetById(id string) (*TaskConfiguration, error) {
	ret, err := s.Repository.GetById(id)

	if err != nil {
		return nil, internalerrors.ErrInternalServerError
	}

	return ret, nil
}
