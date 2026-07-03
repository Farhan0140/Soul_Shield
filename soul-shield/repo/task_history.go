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

type TaskHistoryRepo interface {
	Create(history TaskHistory) error
	GetAll(userID int) ([]TaskHistory, error)
	GetByTaskID(taskID int, userID int,) ([]TaskHistory, error)
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

func (r *taskHistoryRepo) Create(history TaskHistory) error {

	query := `
		INSERT INTO task_histories(
			task_id,
			user_id,
			status,
			completed_at
		) VALUES (
			$1,
			$2,
			$3,
			$4
		)
	`

	_, err := r.db.Exec(
		query,
		history.TaskID,
		history.UserID,
		history.Status,
		history.CompletedAt,
	)

	return err
}

func (r *taskHistoryRepo) GetAll(userID int) ([]TaskHistory, error) {

	var histories []TaskHistory

	query := `
		SELECT
			id,
			task_id,
			user_id,
			status,
			completed_at,
			created_at
		FROM task_histories
		WHERE user_id=$1
		ORDER BY completed_at DESC
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

func (r *taskHistoryRepo) GetByTaskID(taskID int, userID int) ([]TaskHistory, error) {

	var histories []TaskHistory

	query := `
		SELECT
			id,
			task_id,
			user_id,
			status,
			completed_at,
			created_at
		FROM task_histories
		WHERE
			task_id=$1
			AND
			user_id=$2
		ORDER BY completed_at DESC
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