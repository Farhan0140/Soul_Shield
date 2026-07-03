-- +migrate Up

CREATE TABLE task_histories(
    id BIGSERIAL PRIMARY KEY,
    task_id BIGINT NOT NULL,
    user_id BIGINT NOT NULL,
    completed_date DATE NOT NULL,
    completed_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    duration_minutes INTEGER NOT NULL DEFAULT 0 CHECK(duration_minutes>=0),
    note TEXT DEFAULT '',
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    CONSTRAINT fk_history_task
        FOREIGN KEY(task_id)
        REFERENCES tasks(id)
        ON DELETE CASCADE,
        
    CONSTRAINT fk_history_user
        FOREIGN KEY(user_id)
        REFERENCES users(id)
        ON DELETE CASCADE,

    CONSTRAINT unique_task_completion_per_day
        UNIQUE(task_id,completed_date)
);

CREATE INDEX idx_histories_task ON task_histories(task_id);
CREATE INDEX idx_histories_user ON task_histories(user_id);
CREATE INDEX idx_histories_completed_date ON task_histories(completed_date);
CREATE INDEX idx_histories_completed_at ON task_histories(completed_at);