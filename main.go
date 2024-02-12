package main

import (
	"net/http"
	"task-list/internal/domain/infrastructure/database"
	"task-list/internal/domain/taskconfiguration"
	"task-list/internal/endpoints"

	"github.com/go-chi/chi/v5"
)

func main() {

	r := chi.NewRouter()
	data := database.NewDb()
	taskconfigurationRepository := &database.TaskConfigurationRepository{Db: data}
	var service taskconfiguration.Service = taskconfiguration.ServiceImp{Repository: taskconfigurationRepository}
	handler := endpoints.Handler{TaskConfigurationService: service}

	r.Post("/task-configuration", endpoints.HandlerError(handler.TaskConfigurationPost))
	r.Get("/task-configuration", endpoints.HandlerError(handler.TaskConfigurationGet))
	r.Get("/task-configuration/{id}", endpoints.HandlerError(handler.TaskConfigurationGetById))
	http.ListenAndServe(":3000", r)

}
