package middleware

import (
	"context"
	"encoding/json"
	"go-chat-app/internal/entity"
	"go-chat-app/internal/user"
	"go-chat-app/internal/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	UserData userDataKey = "user_data"
)

type (
	UserCtxReq struct {
		UserID   int64  `json:"user_id"`
		Username string `json:"username"`
	}

	MiddleWareImpl struct {
		UserRepo user.Repository
	}

	MiddleWare interface {
		AuthUser() gin.HandlerFunc
	}

	userDataKey string
)

func NewMiddleWare(impl MiddleWareImpl) MiddleWare {
	return &impl
}

func (m *MiddleWareImpl) AuthUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		errResponse := entity.DefaultResponse{Code: http.StatusUnauthorized, Data: struct{}{}}

		ctx := c.Request.Context()

		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			errResponse.Message = "Missing Authorization header"
			c.JSON(http.StatusUnauthorized, errResponse)
			c.Abort()
			return
		}

		splitHeader := strings.Split(authHeader, " ")
		if len(splitHeader) != 2 {
			errResponse.Message = "Invalid Authorization header format"
			c.JSON(http.StatusUnauthorized, errResponse)
			c.Abort()
			return
		}

		if splitHeader[0] != "Bearer" {
			errResponse.Message = "Invalid Authorization scheme"
			c.JSON(http.StatusUnauthorized, errResponse)
			c.Abort()
			return
		}

		token, err := utils.Verify(splitHeader[1])
		if err != nil {
			errResponse.Message = "Unauthorized"
			errResponse.Error = err.Error()
			c.JSON(http.StatusUnauthorized, errResponse)
			c.Abort()
			return
		}

		err = token.Valid()
		if err != nil {
			errResponse.Message = "Unauthorized"
			errResponse.Error = err.Error()
			c.JSON(http.StatusUnauthorized, errResponse)
			c.Abort()
			return
		}

		var userCtx UserCtxReq
		bt, err := json.Marshal(token.Data)
		if err != nil {
			errResponse.Message = "Unauthorized"
			errResponse.Error = err.Error()
			c.JSON(http.StatusUnauthorized, errResponse)
			c.Abort()
			return
		}

		err = json.Unmarshal(bt, &userCtx)
		if err != nil {
			errResponse.Message = "Unauthorized"
			errResponse.Error = err.Error()
			c.JSON(http.StatusUnauthorized, errResponse)
			c.Abort()
			return
		}

		if userCtx.UserID == 0 {
			errResponse.Message = "Unauthorized"
			errResponse.Error = "Invalid user id"
			c.JSON(http.StatusUnauthorized, errResponse)
			c.Abort()
			return
		}

		userData, errGetUser := m.UserRepo.GetUserByUsername(ctx, userCtx.Username)
		if errGetUser != nil {
			errResponse.Message = "Unauthorized"
			errResponse.Error = "Invalid role"
			c.JSON(http.StatusUnauthorized, errResponse)
			c.Abort()
			return
		}
		{
			userCtx.UserID = int64(userData.UserID)
			userCtx.Username = userData.Username
		}

		ctx = context.WithValue(ctx, UserData, userCtx)

		c.Request = c.Request.WithContext(ctx)

		c.Next()
	}
}
