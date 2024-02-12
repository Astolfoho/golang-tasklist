package endpoints

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"task-list/internal/domain/taskconfiguration"
	"task-list/internal/domain/taskconfiguration/contract"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type serviceMock struct {
	mock.Mock
}

func (s *serviceMock) Create(newTaskConfig *contract.NewTaskConfiguration) (string, error) {
	args := s.Called(newTaskConfig)
	return args.String(0), args.Error(1)
}

func (s *serviceMock) GetAll() ([]taskconfiguration.TaskConfiguration, error) {
	args := s.Called()
	return args.Get(0).([]taskconfiguration.TaskConfiguration), args.Error(1)
}

func (s *serviceMock) GetById(id string) (*taskconfiguration.TaskConfiguration, error) {
	args := s.Called()
	return args.Get(0).(*taskconfiguration.TaskConfiguration), args.Error(1)
}

func Test_TaskConfiguration_should_save_new(t *testing.T) {
	assert := assert.New(t)
	service := new(serviceMock)
	body := contract.NewTaskConfiguration{
		Name:        "test_1",
		Description: "description task configuration works",
		DaysOfWeek:  6,
	}
	service.On("Create", mock.MatchedBy(func(r *contract.NewTaskConfiguration) bool {
		if body.Name != r.Name {
			return false
		}

		if body.Description != r.Description {
			return false
		}

		if body.DaysOfWeek != r.DaysOfWeek {
			return false
		}

		return true
	})).Return("TEST_ID", nil)
	handler := Handler{
		TaskConfigurationService: service,
	}

	var buf bytes.Buffer
	json.NewEncoder(&buf).Encode(body)
	req, _ := http.NewRequest("POST", "/task-configuration", &buf)
	rr := httptest.NewRecorder()

	js, status, err := handler.TaskConfigurationPost(rr, req)
	assert.Equal(201, status)
	assert.Nil(err)
	assert.Equal(js.(map[string]string)["id"], "TEST_ID")
}

func Test_TaskConfiguration_should_decodejsonError(t *testing.T) {
	assert := assert.New(t)
	service := new(serviceMock)
	body := "####"
	handler := Handler{
		TaskConfigurationService: service,
	}
	var buf bytes.Buffer
	json.NewEncoder(&buf).Encode(body)
	req, _ := http.NewRequest("POST", "/task-configuration", &buf)
	rr := httptest.NewRecorder()

	_, status, err := handler.TaskConfigurationPost(rr, req)
	assert.Equal(400, status)
	assert.NotNil(err)
}
