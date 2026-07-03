package repo

import (
	"github.com/jmoiron/sqlx"
)

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