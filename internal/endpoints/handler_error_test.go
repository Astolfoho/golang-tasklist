package endpoints

import (
	"net/http"
	"net/http/httptest"
	internalerrors "task-list/internal/domain/internal-errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Handler_Error_When_Endpoints_Return_Bad_Request(t *testing.T) {
	assert := assert.New(t)

	endpoint := func(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
		return nil, 200, internalerrors.ErrInternalServerError
	}

	hadlerError := HandlerError(endpoint)
	req, _ := http.NewRequest("GET", "/", nil)
	res := httptest.NewRecorder()
	hadlerError.ServeHTTP(res, req)

	assert.Equal(http.StatusBadRequest, res.Code)
}

func Test_Handler_Error_When_Endpoints_Return_Sucess(t *testing.T) {
	assert := assert.New(t)

	endpoint := func(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
		return nil, 200, nil
	}

	hadlerError := HandlerError(endpoint)
	req, _ := http.NewRequest("GET", "/", nil)
	res := httptest.NewRecorder()
	hadlerError.ServeHTTP(res, req)

	assert.Equal(http.StatusOK, res.Code)
}
