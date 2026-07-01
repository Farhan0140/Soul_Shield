package task

import (
	"net/http"

	"soulsheld/util"
)

// GetTasks godoc
//
// @Summary Get Tasks
// @Description Get all tasks for current user
// @Tags Tasks
// @Produce json
// @Security BearerAuth
// @Success 200 {array} repo.Task
// @Router /tasks [get]
func (h *Handler) GetTasks(
	w http.ResponseWriter,
	r *http.Request,
) {

	userID, ok := r.Context().Value("userID").(int)

	if !ok {
		util.SendError(
			w,
			map[string]string{
				"error": "Unauthorized Fuck You",
			},
			http.StatusUnauthorized,
		)
		return
	}

	tasks, err := h.taskRepo.GetAll(
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
		tasks,
		http.StatusOK,
	)
}
