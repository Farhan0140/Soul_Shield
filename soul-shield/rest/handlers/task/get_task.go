package task

import (
	"errors"
	"net/http"
	"strconv"

	"soulsheld/util"
)

// GetTask godoc
//
// @Summary Get Task By ID
// @Description Get a single task by its ID
// @Tags Tasks
// @Produce json
// @Security BearerAuth
// @Param id path int true "Task ID"
// @Success 200 {object} repo.Task
// @Failure 400 {object} ErrorResponse
// @Failure 401 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /tasks/{id} [get]
func (h *Handler) GetTask(
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

	userID, ok := r.Context().
		Value("userID").
		(int)

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

	task, err := h.taskRepo.GetByID(
		id,
		userID,
	)

	if err != nil {

		if errors.Is(
			err,
			util.ErrTaskNotFound,
		) {

			util.SendError(
				w,
				map[string]string{
					"error": "task not found",
				},
				http.StatusNotFound,
			)

			return
		}

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
		http.StatusOK,
	)
}