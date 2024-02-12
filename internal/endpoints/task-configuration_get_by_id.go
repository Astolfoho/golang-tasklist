package endpoints

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

func (h *Handler) TaskConfigurationGetById(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
	render.Status(r, 201)
	id := chi.URLParam(r, "id")
	ret, err := h.TaskConfigurationService.GetById(id)
	return ret, 200, err
}
