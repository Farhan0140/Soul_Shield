package repo

import (
	"database/sql"
	"soulsheld/util"
	"time"

	"github.com/jmoiron/sqlx"
)

type Task struct {
	ID          int `json:"id" db:"id"`
	UserID      int `json:"user_id" db:"user_id"`
	Title       string `json:"title" db:"title"`
	Description string `json:"description" db:"description"`
	Priority    string `json:"priority" db:"priority"`
	Status      string `json:"status" db:"status"`
	DueDate     time.Time `json:"due_date" db:"due_date"`
}

type TaskRepo interface {
	Create(task Task) (*Task, error)
	GetAll(userID int) ([]Task, error)
	GetByID(id int, userID int) (*Task, error)
	Update(task Task) error
	Delete(id int, userID int) error
	Complete(id int, userID int) error
}

type taskRepo struct {
	db *sqlx.DB
}

func NewTaskRepo(db *sqlx.DB) TaskRepo {
	return &taskRepo{
		db: db,
	}
}

func (r *taskRepo) Create(task Task) (*Task, error) {

	query := `
		INSERT INTO tasks (
			user_id,
			title,
			description,
			priority,
			status,
			due_date
		)
		VALUES (
			$1,$2,$3,$4,$5,$6
		)
		RETURNING id
	`

	err := r.db.QueryRow(
		query,
		task.UserID,
		task.Title,
		task.Description,
		task.Priority,
		task.Status,
		task.DueDate,
	).Scan(&task.ID)

	if err != nil {
		return nil, err
	}

	return &task, nil
}

func (r *taskRepo) GetAll(userID int) ([]Task, error) {
	var tasks []Task

	query := `
		SELECT
			id,
			user_id,
			title,
			description,
			priority,
			status,
			due_date
		FROM tasks
		WHERE user_id = $1
		ORDER BY created_at DESC
	`

	err := r.db.Select(
		&tasks,
		query,
		userID,
	)

	if err != nil {
		return nil, err
	}

	return tasks, nil
}

func (r *taskRepo) GetByID(id int, userID int) (*Task, error) {

	var task Task

	query := `
		SELECT
			id,
			user_id,
			title,
			description,
			priority,
			status,
			due_date
		FROM tasks
		WHERE id = $1
		AND user_id = $2
	`

	err := r.db.Get(
		&task,
		query,
		id,
		userID,
	)

	if err != nil {

		if err == sql.ErrNoRows {
			return nil, util.ErrTaskNotFound
		}

		return nil, err
	}

	return &task, nil
}

func (r *taskRepo) Update(task Task) error {

	result, err := r.db.Exec(`
		UPDATE tasks
		SET
			title = $1,
			description = $2,
			priority = $3,
			due_date = $4,
			updated_at = NOW()
		WHERE id = $5
		AND user_id = $6
	`,
		task.Title,
		task.Description,
		task.Priority,
		task.DueDate,
		task.ID,
		task.UserID,
	)

	if err != nil {
		return err
	}

	rowsAffected, _ := result.RowsAffected()

	if rowsAffected == 0 {
		return util.ErrTaskNotFound
	}

	return nil
}

func (r *taskRepo) Delete(id int, userID int) error {

	result, err := r.db.Exec(`
		DELETE
		FROM tasks
		WHERE id = $1
		AND user_id = $2
	`,
		id,
		userID,
	)

	if err != nil {
		return err
	}

	rowsAffected, _ := result.RowsAffected()

	if rowsAffected == 0 {
		return util.ErrTaskNotFound
	}

	return nil
}

func (r *taskRepo) Complete(id int, userID int) error {

	tx, err := r.db.Beginx()
	if err != nil {
		return err
	}

	defer tx.Rollback()

	// Update task status
	result, err := tx.Exec(`
		UPDATE tasks
		SET
			status = 'completed',
			updated_at = NOW()
		WHERE
			id = $1
			AND user_id = $2
	`,
		id,
		userID,
	)

	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return util.ErrTaskNotFound
	}

	// Insert history
	_, err = tx.Exec(`
		INSERT INTO task_histories (
			task_id,
			user_id,
			status,
			completed_at
		)
		VALUES (
			$1,
			$2,
			$3,
			NOW()
		)
	`,
		id,
		userID,
		util.TaskCompleted,
	)

	if err != nil {
		return err
	}

	// Commit transaction
	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}