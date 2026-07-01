package task

import (
	"net/http"
	"soulsheld/rest/middlewares"
)

func (h *Handler) RegisterRoutes(mux *http.ServeMux, manager *middlewares.Manager) {
	mux.Handle(
		"POST /tasks",

		manager.With(
			http.HandlerFunc(h.CreateTask),
			h.middlewares.AuthenticateJWT,
		),
	)

	mux.Handle(
		"GET /tasks",

		manager.With(
			http.HandlerFunc(h.GetTasks),
			h.middlewares.AuthenticateJWT,
		),
	)

	mux.Handle(
		"GET /tasks/{id}",

		manager.With(
			http.HandlerFunc(h.GetTask),
			h.middlewares.AuthenticateJWT,
		),
	)

	mux.Handle(
		"PUT /tasks/{id}",

		manager.With(
			http.HandlerFunc(h.UpdateTask),
			h.middlewares.AuthenticateJWT,
		),
	)

	mux.Handle(
		"DELETE /tasks/{id}",

		manager.With(
			http.HandlerFunc(h.DeleteTask),
			h.middlewares.AuthenticateJWT,
		),
	)

	mux.Handle(
		"PATCH /tasks/{id}/complete",

		manager.With(
			http.HandlerFunc(h.CompleteTask),
			h.middlewares.AuthenticateJWT,
		),
	)
}