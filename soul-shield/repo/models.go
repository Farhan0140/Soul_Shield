package repo

import "time"

// User
type User struct {
	ID        int64  `json:"id" db:"id"`
	Full_Name string `json:"full_name" db:"full_name"`
	Email     string `json:"email" db:"email"`
	Password  string `json:"password" db:"password"`
}

// Task
type Task struct {
	ID                int64      `db:"id" json:"id"`
	UserID            int64      `db:"user_id" json:"user_id"`
	CategoryID        *int64     `db:"category_id" json:"category_id,omitempty"`
	Title             string     `db:"title" json:"title"`
	Description       string     `db:"description" json:"description"`
	Priority          string     `db:"priority" json:"priority"`
	RepeatType        string     `db:"repeat_type" json:"repeat_type"`
	RepeatInterval    int        `db:"repeat_interval" json:"repeat_interval"`
	RepeatDays        *string    `db:"repeat_days" json:"repeat_days,omitempty"`
	StartDate         time.Time  `db:"start_date" json:"start_date"`
	EndDate           *time.Time `db:"end_date" json:"end_date,omitempty"`
	StartTime         *string    `db:"start_time" json:"start_time,omitempty"`
	EstimatedMinutes  int        `db:"estimated_minutes" json:"estimated_minutes"`
	LastCompletedDate *time.Time `db:"last_completed_date" json:"last_completed_date,omitempty"`
	NextDueDate       time.Time  `db:"next_due_date" json:"next_due_date"`
	Status            string     `db:"status" json:"status"`
	CreatedAt         time.Time  `db:"created_at" json:"created_at"`
	UpdatedAt         time.Time  `db:"updated_at" json:"updated_at"`
}

type TaskFilter struct {
	Page int
	Limit int
	Search string
	CategoryID *int64
	Priority *string
	RepeatType *string
	Status *string
}

// Task History
type TaskHistory struct {
	ID          int64     `db:"id" json:"id"`
	TaskID      int64     `db:"task_id" json:"task_id"`
	UserID      int64     `db:"user_id" json:"user_id"`
	Status      string    `db:"status" json:"status"`
	CompletedAt time.Time `db:"completed_at" json:"completed_at"`
	CreatedAt   time.Time `db:"created_at" json:"created_at"`
}

type TaskHistoryWithTask struct {
	ID          int64     `db:"id" json:"id"`
	TaskID      int64     `db:"task_id" json:"task_id"`
	TaskTitle   string    `db:"task_title" json:"task_title"`
	Status      string    `db:"status" json:"status"`
	CompletedAt time.Time `db:"completed_at" json:"completed_at"`
}

// Category
type Category struct {
	ID        int64     `db:"id" json:"id"`
	UserID    int64     `db:"user_id" json:"user_id"`
	Name      string    `db:"name" json:"name"`
	Color     string    `db:"color" json:"color"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}
