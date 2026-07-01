package task

import (
	"soulsheld/config"
	"soulsheld/repo"
	"soulsheld/rest/middlewares"
)

type Handler struct {
	cnf         *config.Config
	taskRepo    repo.TaskRepo
	middlewares *middlewares.Middleware
}

func NewHandler(
	cnf *config.Config,
	taskRepo repo.TaskRepo,
	middlewares *middlewares.Middleware,
) *Handler {
	return &Handler{
		cnf:         cnf,
		taskRepo: taskRepo,
		middlewares: middlewares,
	}
}
