-- +migrate Up

CREATE TABLE tasks(
    id BIGSERIAL PRIMARY KEY,
    user_id BIGINT NOT NULL,
    category_id BIGINT,
    title VARCHAR(255) NOT NULL CHECK(length(trim(title))>0),
    description TEXT DEFAULT '',
    priority VARCHAR(20) NOT NULL DEFAULT 'medium' CHECK(priority IN('low','medium','high')),
    repeat_type VARCHAR(20) NOT NULL DEFAULT 'none' CHECK(repeat_type IN('none','daily','weekly','monthly')),
    start_date DATE NOT NULL DEFAULT CURRENT_DATE,
    end_date DATE,
    start_time TIME,
    estimated_minutes INTEGER NOT NULL DEFAULT 0 CHECK(estimated_minutes>=0),
    is_active BOOLEAN NOT NULL DEFAULT TRUE,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    
    CONSTRAINT fk_task_user
        FOREIGN KEY(user_id)
        REFERENCES users(id)
        ON DELETE CASCADE,
    CONSTRAINT fk_task_category
        FOREIGN KEY(category_id)
        REFERENCES categories(id)
        ON DELETE SET NULL
);

CREATE INDEX idx_tasks_user ON tasks(user_id);
CREATE INDEX idx_tasks_category ON tasks(category_id);
CREATE INDEX idx_tasks_repeat_type ON tasks(repeat_type);
CREATE INDEX idx_tasks_active ON tasks(is_active);
CREATE INDEX idx_tasks_start_date ON tasks(start_date);