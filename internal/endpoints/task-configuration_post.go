package endpoints

import (
	"net/http"
	"task-list/internal/domain/taskconfiguration/contract"

	"github.com/go-chi/render"
)

func (h *Handler) TaskConfigurationPost(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
	var request contract.NewTaskConfiguration
	err := render.DecodeJSON(r.Body, &request)

	if err != nil {
		return nil, 400, err
	}

	id, err := h.TaskConfigurationService.Create(&request)
	return map[string]string{"id": id}, 201, err
}
