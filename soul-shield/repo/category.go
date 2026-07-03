package repo

import (
	"database/sql"
	"errors"
	"soulsheld/util"

	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

type CategoryRepo interface {
	Create(category Category) (*Category, error)
	GetAll(userID int) ([]Category, error)
	GetByID(id, userID int) (*Category, error)
	Update(category Category) error
	Delete(id, userID int) error
}

type categoryRepo struct {
	db *sqlx.DB
}

func NewCategoryRepo(db *sqlx.DB) CategoryRepo {
	return &categoryRepo{
		db: db,
	}
}

func (r *categoryRepo) Create(category Category) (*Category, error) {

	query := `
		INSERT INTO categories(
			user_id,
			name,
			color
		)
		VALUES(
			$1,
			$2,
			$3
		)
		RETURNING id,created_at,updated_at
	`

	err := r.db.QueryRow(
		query,
		category.UserID,
		category.Name,
		category.Color,
	).Scan(
		&category.ID,
		&category.CreatedAt,
		&category.UpdatedAt,
	)

	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			if pqErr.Code == "23505" {
				return nil, util.ErrCategoryExists
			}
		}
		return nil, err
	}

	return &category, nil
}

func (r *categoryRepo) GetAll(userID int) ([]Category, error) {

	var categories []Category

	query := `
		SELECT
			id,
			user_id,
			name,
			color,
			created_at,
			updated_at
		FROM categories
		WHERE user_id=$1
		ORDER BY name ASC
	`

	err := r.db.Select(
		&categories,
		query,
		userID,
	)

	if err != nil {
		return nil, err
	}

	return categories, nil
}

func (r *categoryRepo) GetByID(id, userID int) (*Category, error) {

	var category Category

	query := `
		SELECT
			id,
			user_id,
			name,
			color,
			created_at,
			updated_at
		FROM categories
		WHERE
			id=$1
		AND
			user_id=$2
		LIMIT 1
	`

	err := r.db.Get(
		&category,
		query,
		id,
		userID,
	)

	if err != nil {

		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}

		return nil, err
	}

	return &category, nil
}

func (r *categoryRepo) Update(category Category) error {

	result, err := r.db.Exec(`
		UPDATE categories
		SET
			name=$1,
			color=$2,
			updated_at=NOW()
		WHERE
			id=$3
		AND
			user_id=$4
	`,
		category.Name,
		category.Color,
		category.ID,
		category.UserID,
	)

	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return sql.ErrNoRows
	}

	return nil
}

func (r *categoryRepo) Delete(id, userID int) error {

	result, err := r.db.Exec(`
		DELETE FROM categories
		WHERE
			id=$1
		AND
			user_id=$2
	`,
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
		return sql.ErrNoRows
	}

	return nil
}
