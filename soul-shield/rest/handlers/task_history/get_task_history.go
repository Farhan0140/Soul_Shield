package task_history

import (
	"net/http"
	"strconv"

	"soulsheld/util"
)

// GetTaskHistory godoc
//
// @Summary Get History Of One Task
// @Description Get all completion history of a specific task
// @Tags Task History
// @Produce json
// @Security BearerAuth
// @Param id path int true "Task ID"
// @Success 200 {array} TaskHistoryResponse
// @Failure 400 {object} ErrorResponse
// @Failure 401 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /tasks/{id}/history [get]
func (h *Handler) GetTaskHistory(
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

	history, err := h.taskHistoryRepo.GetByTaskID(
		id,
		userID,
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
		history,
		http.StatusOK,
	)
}