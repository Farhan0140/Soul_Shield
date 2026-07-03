package task_history

import (
	"net/http"
	"soulsheld/rest/middlewares"
)

func (h *Handler) RegisterRoutes(mux *http.ServeMux, manager *middlewares.Manager) {
	mux.Handle(
		"GET /tasks/history",
		manager.With(
			http.HandlerFunc(h.GetHistory),
			h.middlewares.AuthenticateJWT,
		),
	)

	mux.Handle(
		"GET /tasks/{id}/history",

		manager.With(
			http.HandlerFunc(h.GetTaskHistory),
			h.middlewares.AuthenticateJWT,
		),
	)
}