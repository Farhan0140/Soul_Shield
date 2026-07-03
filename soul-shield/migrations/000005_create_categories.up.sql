-- +migrate Up
CREATE TABLE categories(
    id BIGSERIAL PRIMARY KEY,
    user_id BIGINT NOT NULL,
    name VARCHAR(100) NOT NULL CHECK(length(trim(name))>0),
    color VARCHAR(20) NOT NULL DEFAULT '#3B82F6' CHECK(color ~ '^#[0-9A-Fa-f]{6}$'),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    CONSTRAINT fk_category_user
    FOREIGN KEY(user_id)
    REFERENCES users(id)
    ON DELETE CASCADE,
    CONSTRAINT unique_user_category
    UNIQUE(user_id,name)
);

CREATE INDEX idx_categories_user ON categories(user_id);