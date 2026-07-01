package task

import (
	"net/http"
	"strconv"

	"soulsheld/util"
)

// DeleteTask godoc
//
// @Summary Delete Task
// @Description Delete a task
// @Tags Tasks
// @Produce json
// @Security BearerAuth
// @Param id path int true "Task ID"
// @Success 200 {object} SuccessResponse
// @Router /tasks/{id} [delete]
func (h *Handler) DeleteTask(
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

	err = h.taskRepo.Delete(
		id,
		userID,
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
			"message": "task deleted",
		},
		http.StatusOK,
	)
}