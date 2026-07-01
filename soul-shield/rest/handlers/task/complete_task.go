package task

import (
	"net/http"
	"strconv"

	"soulsheld/util"
)

// CompleteTask godoc
//
// @Summary Complete Task
// @Description Mark task as completed
// @Tags Tasks
// @Produce json
// @Security BearerAuth
// @Param id path int true "Task ID"
// @Success 200 {object} SuccessResponse
// @Router /tasks/{id}/complete [patch]
func (h *Handler) CompleteTask(
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

	err = h.taskRepo.Complete(
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
			"message": "task completed",
		},
		http.StatusOK,
	)
}