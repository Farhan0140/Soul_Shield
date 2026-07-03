-- +migrate Down
DROP INDEX IF EXISTS idx_categories_user;
DROP TABLE IF EXISTS categories;