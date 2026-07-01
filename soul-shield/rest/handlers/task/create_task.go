package task

import (
	"encoding/json"
	"net/http"
	"time"

	"soulsheld/repo"
	"soulsheld/util"
)

// CreateTask godoc
//
// @Summary Create Task
// @Description Create a new task
// @Tags Tasks
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param body body CreateTaskRequest true "Task Info"
// @Success 201 {object} repo.Task
// @Failure 400 {object} ErrorResponse
// @Router /tasks [post]
func (h *Handler) CreateTask(
	w http.ResponseWriter,
	r *http.Request,
) {

	var req CreateTaskRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		util.SendError(
			w,
			map[string]string{
				"error": "invalid request body",
			},
			http.StatusBadRequest,
		)
		return
	}

	userID, ok := r.Context().Value("userID").(int)

	if !ok {
		util.SendError(
			w,
			map[string]string{
				"error": "unauthorized",
			},
			http.StatusUnauthorized,
		)
		return
	}

	dueDate, err := time.Parse(
		time.RFC3339,
		req.DueDate,
	)

	if err != nil {
		util.SendError(
			w,
			map[string]string{
				"error": "invalid due_date",
			},
			http.StatusBadRequest,
		)
		return
	}

	task, err := h.taskRepo.Create(
		repo.Task{
			UserID:      userID,
			Title:       req.Title,
			Description: req.Description,
			Priority:    req.Priority,
			Status:      "pending",
			DueDate:     dueDate,
		},
	)

	if err != nil {
		util.SendError(
			w,
			map[string]string{
				"error": err.Error(),
			},
			http.StatusInternalServerError,
		)
		return
	}

	util.SendData(
		w,
		task,
		http.StatusCreated,
	)
}