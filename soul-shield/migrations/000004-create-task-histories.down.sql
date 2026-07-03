-- +migrate Down

DROP INDEX IF EXISTS idx_histories_completed_at;
DROP INDEX IF EXISTS idx_histories_completed_date;
DROP INDEX IF EXISTS idx_histories_user;
DROP INDEX IF EXISTS idx_histories_task;

DROP TABLE IF EXISTS task_histories;