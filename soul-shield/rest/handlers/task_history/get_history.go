package task_history

import (
	"net/http"

	"soulsheld/util"
)

// GetHistory godoc
//
// @Summary Get Task History
// @Description Get all completed task history of authenticated user
// @Tags Task History
// @Produce json
// @Security BearerAuth
// @Success 200 {array} HistoryResponse
// @Failure 401 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /tasks/history [get]
func (h *Handler) GetHistory(w http.ResponseWriter, r *http.Request) {

	userID, ok := r.Context().
		Value("userID").
		(int)

	if !ok {

		util.SendError(
			w,
			map[string]string{
				"error": "Unauthorized",
			},
			http.StatusUnauthorized,
		)

		return
	}

	history, err := h.taskHistoryRepo.GetAll(userID)

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