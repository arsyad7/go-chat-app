package user

import (
	"context"
	"database/sql"
	"log/slog"
)

type DBTX interface {
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
	PrepareContext(context.Context, string) (*sql.Stmt, error)
	QueryContext(context.Context, string, ...interface{}) (*sql.Rows, error)
	QueryRowContext(context.Context, string, ...interface{}) *sql.Row
}

type repository struct {
	db DBTX
}

func NewRepository(db DBTX) Repository {
	return &repository{db: db}
}

func (r *repository) CreateUser(ctx context.Context, user User) (id int64, err error) {
	err = r.db.QueryRowContext(ctx, "INSERT INTO users (username, password) VALUES ($1, $2) RETURNING user_id", user.Username, user.Password).Scan(&id)
	if err != nil {
		slog.Error("Error inserting user: %v", err)
		return 0, err
	}

	return id, err
}

func (r *repository) GetUserByUsername(ctx context.Context, username string) (user User, err error) {
	row := r.db.QueryRowContext(ctx, "SELECT user_id, username, password FROM users WHERE username=$1", username)
	err = row.Scan(&user.UserID, &user.Username, &user.Password)
	if err != nil {
		return user, err
	}
	return user, nil
}
