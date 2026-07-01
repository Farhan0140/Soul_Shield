package task

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"soulsheld/repo"
	"soulsheld/util"
)

// UpdateTask godoc
//
// @Summary Update Task
// @Description Update existing task
// @Tags Tasks
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path int true "Task ID"
// @Param body body UpdateTaskRequest true "Task Info"
// @Success 200 {object} SuccessResponse
// @Router /tasks/{id} [put]
func (h *Handler) UpdateTask(
	w http.ResponseWriter,
	r *http.Request,
) {

	id, err := strconv.Atoi(
		r.PathValue("id"),
	)

	if err != nil {
		util.SendError(
			w,
			map[string]string{
				"error": "invalid task id",
			},
			http.StatusBadRequest,
		)
		return
	}

	userID := r.Context().
		Value("userID").
		(int)

	var req UpdateTaskRequest

	if err := json.NewDecoder(
		r.Body,
	).Decode(
		&req,
	); err != nil {

		util.SendError(
			w,
			map[string]string{
				"error": "invalid request body",
			},
			http.StatusBadRequest,
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
				"error": "invalid due date",
			},
			http.StatusBadRequest,
		)
		return
	}

	err = h.taskRepo.Update(
		repo.Task{
			ID:          id,
			UserID:      userID,
			Title:       req.Title,
			Description: req.Description,
			Priority:    req.Priority,
			DueDate:     dueDate,
		},
	)

	if err != nil {

		util.SendError(
			w,
			map[string]string{
				"error": err.Error(),
			},
			http.StatusBadRequest,
		)

		return
	}

	util.SendData(
		w,
		map[string]string{
			"message": "task updated",
		},
		http.StatusOK,
	)
}