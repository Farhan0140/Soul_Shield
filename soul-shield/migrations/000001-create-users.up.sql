
-- +migrate Up
CREATE TABLE users(
	id BIGSERIAL PRIMARY KEY,
	full_name VARCHAR(100) NOT NULL CHECK(length(trim(full_name))>0),
	email VARCHAR(255) NOT NULL UNIQUE,
	password TEXT NOT NULL,
	role VARCHAR(20) NOT NULL DEFAULT 'user' CHECK(role IN('user','admin')),
	created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
	updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_users_email ON users(email);