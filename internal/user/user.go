package user

import (
	"context"
	"go-chat-app/internal/entity"
)

type User struct {
	UserID   int64  `json:"user_id,omitempty" db:"user_id"`
	Username string `json:"username" db:"username"`
	Password string `json:"password" db:"password"`
}

type Repository interface {
	CreateUser(ctx context.Context, user User) (id int64, err error)
	GetUserByUsername(ctx context.Context, username string) (user User, err error)
}

type Service interface {
	UserRegistration(ctx context.Context, form User) (resp entity.DefaultResponse, err error)
	UserLogin(ctx context.Context, form User) (resp entity.DefaultResponse, err error)
}
