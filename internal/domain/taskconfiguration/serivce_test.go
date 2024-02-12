package taskconfiguration

import (
	"errors"
	internalerrors "task-list/internal/domain/internal-errors"
	"task-list/internal/domain/taskconfiguration/contract"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	newTaskConfig = contract.NewTaskConfiguration{
		Name:        "Test 1",
		Description: "Test Description 1",
		DaysOfWeek:  int(Sunday | Monday),
	}
)

type repositoryMock struct {
	mock.Mock
}

func (r *repositoryMock) Create(taskConfig *TaskConfiguration) error {
	args := r.Called(taskConfig)
	return args.Error(0)
}

func (r *repositoryMock) GetAll() ([]TaskConfiguration, error) {
	//args := r.Called(taskConfig)
	return nil, nil
}

func (r *repositoryMock) GetById(id string) (*TaskConfiguration, error) {
	//args := r.Called(taskConfig)
	return nil, nil
}

func Test_Service_Create(t *testing.T) {

	repositoryMock := new(repositoryMock)
	repositoryMock.On("Create", mock.MatchedBy(func(tc *TaskConfiguration) bool {

		if tc.Name != newTaskConfig.Name {
			return false
		}

		if tc.Description != newTaskConfig.Description {
			return false
		}

		if int(tc.DaysOfWeek) != newTaskConfig.DaysOfWeek {
			return false
		}

		return true
	})).Return(nil)
	service := ServiceImp{Repository: repositoryMock}
	service.Create(&newTaskConfig)

	repositoryMock.AssertExpectations(t)
}

func Test_Service_ValidateDomainError(t *testing.T) {
	assert := assert.New(t)
	newTaskConfig.Name = ""
	repositoryMock := new(repositoryMock)

	service := ServiceImp{Repository: repositoryMock}

	_, err := service.Create(&newTaskConfig)
	assert.NotNil(err)
}

func Test_Service_ValidateDatabaseError(t *testing.T) {
	assert := assert.New(t)
	repositoryMock := new(repositoryMock)
	repositoryMock.On("Create", mock.Anything).Return(errors.New("error to save on database"))

	service := ServiceImp{Repository: repositoryMock}
	newTaskConfig.Name = "Test 1"
	_, err := service.Create(&newTaskConfig)

	assert.ErrorIs(err, internalerrors.ErrInternalServerError)
}
