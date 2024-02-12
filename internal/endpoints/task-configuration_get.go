package endpoints

import (
	"net/http"

	"github.com/go-chi/render"
)

func (h *Handler) TaskConfigurationGet(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
	render.Status(r, 201)
	ret, err := h.TaskConfigurationService.GetAll()
	return ret, 200, err
}
