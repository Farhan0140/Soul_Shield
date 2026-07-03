package repo

import (
	"time"

	"github.com/jmoiron/sqlx"
)

type TaskHistory struct {
	ID          int       `db:"id" json:"id"`
	TaskID      int       `db:"task_id" json:"task_id"`
	UserID      int       `db:"user_id" json:"user_id"`
	Status      string    `db:"status" json:"status"`
	CompletedAt time.Time `db:"completed_at" json:"completed_at"`
	CreatedAt   time.Time `db:"created_at" json:"created_at"`
}

type TaskHistoryWithTask struct {
	ID          int       `db:"id" json:"id"`
	TaskID      int       `db:"task_id" json:"task_id"`
	TaskTitle   string    `db:"task_title" json:"task_title"`
	Status      string    `db:"status" json:"status"`
	CompletedAt time.Time `db:"completed_at" json:"completed_at"`
}

type TaskHistoryRepo interface {
	GetAll(userID int) ([]TaskHistoryWithTask, error)
	GetByTaskID(taskID int, userID int) ([]TaskHistoryWithTask, error)
}

type taskHistoryRepo struct {
	db *sqlx.DB
}

func NewTaskHistoryRepo(
	db *sqlx.DB,
) TaskHistoryRepo {

	return &taskHistoryRepo{
		db: db,
	}
}

func (r *taskHistoryRepo) GetAll(userID int) ([]TaskHistoryWithTask, error) {

	var histories []TaskHistoryWithTask

	query := `
		SELECT
			th.id,
			th.task_id,
			t.title AS task_title,
			th.status,
			th.completed_at
		FROM task_histories th
		INNER JOIN tasks t
			ON t.id = th.task_id
		WHERE
			th.user_id = $1
		ORDER BY
			th.completed_at DESC
	`

	err := r.db.Select(
		&histories,
		query,
		userID,
	)

	if err != nil {
		return nil, err
	}

	return histories, nil
}

func (r *taskHistoryRepo) GetByTaskID(taskID int, userID int) ([]TaskHistoryWithTask, error) {

	var histories []TaskHistoryWithTask

	query := `
		SELECT
			th.id,
			th.task_id,
			t.title AS task_title,
			th.status,
			th.completed_at
		FROM task_histories th
		INNER JOIN tasks t
			ON t.id = th.task_id
		WHERE
			th.task_id = $1
			AND
			th.user_id = $2
		ORDER BY
			th.completed_at DESC
	`

	err := r.db.Select(
		&histories,
		query,
		taskID,
		userID,
	)

	if err != nil {
		return nil, err
	}

	return histories, nil
}