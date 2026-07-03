package repo

import (
	"database/sql"
	"errors"
	"soulsheld/util"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)


const taskSelectColumns = `
	id,
	user_id,
	category_id,
	title,
	description,
	priority,
	repeat_type,
	repeat_interval,
	repeat_days,
	start_date,
	end_date,
	start_time,
	estimated_minutes,
	last_completed_date,
	next_due_date,
	status,
	created_at,
	updated_at
`

type TaskRepo interface {
	Create(task Task) (*Task, error)

	GetAll(userID int64) ([]Task, error)
	// GetAll(userID int64, filter TaskFilter) ([]Task, error) // For Future
	GetByID(id, userID int64) (*Task, error)

	Update(task Task) error

	Delete(taskId, userID int64) error

	Archive(id, userID int64) error
	Activate(id, userID int64) error

	GetActiveTasks(userID int64, filter TaskFilter) ([]Task, error)
	GetArchivedTasks(userID int64, filter TaskFilter) ([]Task, error)

	GetTodayTasks(userID int64) ([]Task, error)
	// GetTasksByDate(userID int64, date time.Time) ([]Task, error)

	UpdateCompletion(
		taskID int64,
		lastCompleted time.Time,
		nextDueDate time.Time,
	) error
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

	tx, err := r.db.Beginx()
	if err != nil {
		return nil, err
	}

	defer tx.Rollback()

	query := `
		INSERT INTO tasks(
			user_id,
			category_id,
			title,
			description,
			priority,
			repeat_type,
			repeat_interval,
			repeat_days,
			start_date,
			end_date,
			start_time,
			estimated_minutes,
			next_due_date
		)
		VALUES(
			$1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13
		)
		RETURNING
			id,
			created_at,
			updated_at
	`

	err = tx.QueryRow(
		query,
		task.UserID,
		task.CategoryID,
		task.Title,
		task.Description,
		task.Priority,
		task.RepeatType,
		task.RepeatInterval,
		task.RepeatDays,
		task.StartDate,
		task.EndDate,
		task.StartTime,
		task.EstimatedMinutes,
		task.NextDueDate,
	).Scan(
		&task.ID,
		&task.CreatedAt,
		&task.UpdatedAt,
	)

	if err != nil {

		if pqErr, ok := err.(*pq.Error); ok {

			switch pqErr.Code {

			case "23503":
				return nil, util.ErrCategoryNotFound

			case "23514":
				return nil, util.ErrInvalidRequest
			}
		}

		return nil, err
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return &task, nil
}

func (r *taskRepo) GetAll(userID int64) ([]Task, error) {

	var tasks []Task

	query := `
		SELECT
			` + taskSelectColumns + `
		FROM tasks
		WHERE user_id=$1
		ORDER BY
			CASE
				WHEN status='active' THEN 0
			ELSE 1
			END,

			created_at DESC
	`

	err := r.db.Select(
		&tasks,
		query,
		userID,
	)

	if err != nil {
		return nil, err
	}

	if tasks == nil {
		tasks = []Task{}
	}

	return tasks, nil
}

func (r *taskRepo) GetByID(id, userID int64) (*Task, error) {

	var task Task

	query := `
		SELECT
			` + taskSelectColumns + `
		FROM tasks
		WHERE
			id=$1
		AND
			user_id=$2
		LIMIT 1
	`

	err := r.db.Get(
		&task,
		query,
		id,
		userID,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, util.ErrTaskNotFound
		}
		return nil, err
	}

	return &task, nil
}

func (r *taskRepo) Update(task Task) error {

	query := `
		UPDATE tasks
		SET
			category_id=$1,
			title=$2,
			description=$3,
			priority=$4,
			repeat_type=$5,
			repeat_interval=$6,
			repeat_days=$7,
			start_date=$8,
			end_date=$9,
			start_time=$10,
			estimated_minutes=$11,
			next_due_date=$12,
			status=$13,
			updated_at=NOW()
		WHERE
			id=$13
		AND
			user_id=$14
	`

	result, err := r.db.Exec(
		query,
		task.CategoryID,
		task.Title,
		task.Description,
		task.Priority,
		task.RepeatType,
		task.RepeatInterval,
		task.RepeatDays,
		task.StartDate,
		task.EndDate,
		task.StartTime,
		task.EstimatedMinutes,
		task.NextDueDate,
		task.ID,
		task.UserID,
		task.Status,
	)

	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code {

			case "23503":
				return util.ErrCategoryNotFound

			case "23514":
				return util.ErrInvalidRequest
			}
		}

		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return util.ErrTaskNotFound
	}

	return nil
}

func (r *taskRepo) Delete(taskId int64, userID int64) error {

	query := `
		DELETE FROM tasks
		WHERE
			id=$1
		AND
			user_id=$2
	`

	result, err := r.db.Exec(
		query,
		taskId,
		userID,
	)

	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return util.ErrTaskNotFound
	}

	return nil
}

func (r *taskRepo) Archive(id, userID int64) error {

	query := `
		UPDATE tasks
		SET
			status=$1,
			updated_at=NOW()
		WHERE
			id=$2
		AND
			user_id=$3
	`

	result, err := r.db.Exec(
		query,
		util.TaskStatusArchived,
		id,
		userID,
	)

	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return util.ErrTaskNotFound
	}

	return nil
}

func (r *taskRepo) Activate(id, userID int64) error {

	query := `
		UPDATE tasks
		SET
			status=$1,
			updated_at=NOW()
		WHERE
			id=$2
		AND
			user_id=$3
	`

	result, err := r.db.Exec(
		query,
		util.TaskStatusActive,
		id,
		userID,
	)

	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return util.ErrTaskNotFound
	}

	return nil
}

func (r *taskRepo) GetTodayTasks(userID int64) ([]Task, error) {

	var tasks []Task

	query := `
		SELECT
			` + taskSelectColumns + `
		FROM tasks
		WHERE
			user_id=$1
		AND
			status=$2
		AND
			next_due_date=CURRENT_DATE
		ORDER BY
			start_time ASC NULLS LAST,
			CASE priority
				WHEN 'high' THEN 1
				WHEN 'medium' THEN 2
				WHEN 'low' THEN 3
			END,
			created_at ASC
	`

	err := r.db.Select(
		&tasks,
		query,
		userID,
		util.TaskStatusActive,
	)

	if err != nil {
		return nil, err
	}

	if tasks == nil {
		tasks = []Task{}
	}

	return tasks, nil
}

func (r *taskRepo) GetActiveTasks(userID int64, filter TaskFilter) ([]Task, error) {

	var tasks []Task

	query := `
		SELECT
			` + taskSelectColumns + `
		FROM tasks
		WHERE
			user_id=$1
		AND
			status=$2
		ORDER BY
			next_due_date ASC,
			start_time ASC NULLS LAST,
			CASE priority
				WHEN 'high' THEN 1
				WHEN 'medium' THEN 2
				WHEN 'low' THEN 3
			END,
			created_at DESC
	`

	err := r.db.Select(
		&tasks,
		query,
		userID,
		util.TaskStatusActive,
	)

	if err != nil {
		return nil, err
	}

	if tasks == nil {
		tasks = []Task{}
	}

	return tasks, nil
}

func (r *taskRepo) GetArchivedTasks(userID int64, filter TaskFilter) ([]Task, error) {

	var tasks []Task

	query := `
		SELECT
			` + taskSelectColumns + `
		FROM tasks
		WHERE
			user_id=$1
		AND
			status=$2
		ORDER BY
			updated_at DESC
	`

	err := r.db.Select(
		&tasks,
		query,
		userID,
		util.TaskStatusArchived,
	)

	if err != nil {
		return nil, err
	}

	if tasks == nil {
		tasks = []Task{}
	}

	return tasks, nil
}

func (r *taskRepo) UpdateCompletion(taskID int64, lastCompleted time.Time, nextDueDate time.Time) error {

	query := `
		UPDATE tasks
		SET
			last_completed_date=$1,
			next_due_date=$2,
			updated_at=NOW()
		WHERE id=$3
	`

	result, err := r.db.Exec(
		query,
		lastCompleted,
		nextDueDate,
		taskID,
	)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return util.ErrTaskNotFound
	}

	return nil
}