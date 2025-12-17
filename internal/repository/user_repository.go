package repository

import (
	"context"
	"database/sql"
	"time"

	"user-api/internal/models"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) CreateUser(ctx context.Context, name string, dob time.Time) (models.User, error) {
	var user models.User

	query := `
		INSERT INTO users (name, dob)
		VALUES ($1, $2)
		RETURNING id, name, dob
	`

	err := r.db.QueryRowContext(ctx, query, name, dob).
		Scan(&user.ID, &user.Name, &user.Dob)

	return user, err
}

func (r *UserRepository) GetUserByID(ctx context.Context, id int32) (models.User, error) {
	var user models.User

	query := `
		SELECT id, name, dob
		FROM users
		WHERE id = $1
	`

	err := r.db.QueryRowContext(ctx, query, id).
		Scan(&user.ID, &user.Name, &user.Dob)

	return user, err
}

func (r *UserRepository) ListUsers(ctx context.Context) ([]models.User, error) {
	query := `
		SELECT id, name, dob
		FROM users
		ORDER BY id
	`

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User

	for rows.Next() {
		var u models.User
		if err := rows.Scan(&u.ID, &u.Name, &u.Dob); err != nil {
			return nil, err
		}
		users = append(users, u)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

func (r *UserRepository) UpdateUser(ctx context.Context, id int32, name string, dob time.Time) (models.User, error) {
	var user models.User

	query := `
		UPDATE users
		SET name = $1, dob = $2
		WHERE id = $3
		RETURNING id, name, dob
	`

	err := r.db.QueryRowContext(ctx, query, name, dob, id).
		Scan(&user.ID, &user.Name, &user.Dob)

	return user, err
}

func (r *UserRepository) DeleteUser(ctx context.Context, id int32) error {
	query := `DELETE FROM users WHERE id = $1`

	result, err := r.db.ExecContext(ctx, query, id)
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
