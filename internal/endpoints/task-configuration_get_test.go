package endpoints

import (
	"net/http"
	"net/http/httptest"
	"task-list/internal/domain/taskconfiguration"
	"task-list/internal/domain/taskconfiguration/contract"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type serviceMockGet struct {
	mock.Mock
}

func (s *serviceMockGet) Create(newTaskConfig *contract.NewTaskConfiguration) (string, error) {
	args := s.Called(newTaskConfig)
	return args.String(0), args.Error(1)
}

func (s *serviceMockGet) GetAll() ([]taskconfiguration.TaskConfiguration, error) {
	args := s.Called()
	return args.Get(0).([]taskconfiguration.TaskConfiguration), args.Error(1)
}

func (s *serviceMockGet) GetById(id string) (*taskconfiguration.TaskConfiguration, error) {
	args := s.Called()
	return args.Get(0).(*taskconfiguration.TaskConfiguration), args.Error(1)
}

func Test_TaskConfiguration_GetAll(t *testing.T) {
	assert := assert.New(t)
	service := new(serviceMockGet)
	body := taskconfiguration.TaskConfiguration{
		Name:        "test_1",
		Description: "description task configuration works",
		DaysOfWeek:  6,
	}
	list := []taskconfiguration.TaskConfiguration{body}
	service.On("GetAll", mock.Anything).Return(list, nil)
	handler := Handler{
		TaskConfigurationService: service,
	}

	req, _ := http.NewRequest("GET", "/task-configuration", nil)
	rr := httptest.NewRecorder()

	js, status, err := handler.TaskConfigurationGet(rr, req)
	assert.Equal(200, status)
	assert.Nil(err)
	assert.Equal(len(js.([]taskconfiguration.TaskConfiguration)), 1)
}
