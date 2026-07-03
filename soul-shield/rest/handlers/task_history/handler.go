package task_history

import (
	"soulsheld/config"
	"soulsheld/repo"
	"soulsheld/rest/middlewares"
)

type Handler struct {
	cnf             *config.Config
	taskHistoryRepo repo.TaskHistoryRepo
	middlewares     *middlewares.Middleware
}

func NewHandler(
	cnf *config.Config,
	taskHistoryRepo repo.TaskHistoryRepo,
	middlewares *middlewares.Middleware,
) *Handler {

	return &Handler{
		cnf:             cnf,
		taskHistoryRepo: taskHistoryRepo,
		middlewares:     middlewares,
	}
}
