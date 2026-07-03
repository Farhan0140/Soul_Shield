package task_history

import "time"

type HistoryResponse struct {
	ID int `json:"id"`
	TaskID int `json:"task_id"`
	TaskTitle string `json:"task_title"`
	Status string `json:"status"`
	CompletedAt time.Time `json:"completed_at"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

