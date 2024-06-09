package user

import (
	"context"
	"go-chat-app/internal/entity"
	"go-chat-app/internal/utils"
	"log/slog"
	"net/http"
	"time"
)

type (
	JWTData struct {
		Username string `json:"username"`
		UserID   int64  `json:"user_id"`
	}

	service struct {
		Repository
		timeout time.Duration
	}
)

func NewService(repository Repository) Service {
	return &service{
		repository,
		time.Duration(2) * time.Second,
	}
}

func (u *service) UserRegistration(ctx context.Context, form User) (resp entity.DefaultResponse, err error) {
	{
		resp.Code = http.StatusCreated
		resp.Message = "success"
	}

	form.Password, err = utils.HashPassword(form.Password)
	if err != nil {
		slog.Error("Error hashing password: %v", err)
		return resp, err
	}

	userId, err := u.CreateUser(ctx, form)
	if err != nil {
		slog.Error("Error creating user: %v", err)
		return resp, err
	}

	token, exp, err := utils.Sign(JWTData{
		UserID:   int64(userId),
		Username: form.Username,
	})
	if err != nil {
		slog.Error("Error signing token: %v", err)
		return resp, err
	}

	resp.Data = struct {
		Token  string `json:"token"`
		Expire int64  `json:"expire"`
	}{
		Token:  token,
		Expire: exp,
	}

	return resp, nil
}

func (u *service) UserLogin(ctx context.Context, form User) (resp entity.DefaultResponse, err error) {
	{
		resp.Code = http.StatusOK
		resp.Message = "success"
	}

	user, err := u.GetUserByUsername(ctx, form.Username)
	if err != nil {
		slog.Error("Error getting user: %v", err)
		return resp, err
	}

	if user.UserID == 0 {
		resp.Code = http.StatusUnauthorized
		resp.Message = "unauthorized"
		return resp, nil
	}

	if err = utils.VerifyPassword(form.Password, user.Password); err != nil {
		resp.Code = http.StatusUnauthorized
		resp.Message = "unauthorized"
		return resp, nil
	}

	token, exp, err := utils.Sign(JWTData{
		UserID:   user.UserID,
		Username: user.Username,
	})
	if err != nil {
		slog.Error("Error signing token: %v", err)
		return resp, err
	}

	resp.Data = struct {
		Token  string `json:"token"`
		Expire int64  `json:"expire"`
	}{
		Token:  token,
		Expire: exp,
	}

	return resp, nil
}
