-- +migrate Down

DROP INDEX IF EXISTS idx_tasks_start_date;
DROP INDEX IF EXISTS idx_tasks_active;
DROP INDEX IF EXISTS idx_tasks_repeat_type;
DROP INDEX IF EXISTS idx_tasks_category;
DROP INDEX IF EXISTS idx_tasks_user;

DROP TABLE IF EXISTS tasks;