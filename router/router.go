package router

import (
	"go-chat-app/internal/user"
	"go-chat-app/internal/ws"
	"go-chat-app/pkg/middleware"

	"github.com/gin-gonic/gin"
)

var r *gin.Engine

func InitRouter(userHandler *user.Handler, wsHandler *ws.Handler, userRepo user.Repository) {
	r = gin.Default()

	mw := middleware.NewMiddleWare(middleware.MiddleWareImpl{UserRepo: userRepo})

	r.POST("/signup", userHandler.CreateUser)
	r.POST("/login", userHandler.Login)

	authGroup := r.Group("/ws")
	authGroup.Use(mw.AuthUser())
	{
		authGroup.POST("/createRoom", wsHandler.CreateRoom)
		authGroup.GET("/joinRoom/:roomId", wsHandler.JoinRoom)
		authGroup.GET("/getRooms", wsHandler.GetRooms)
		authGroup.GET("/getClients/:roomId", wsHandler.GetClients)
	}
}

func Start(addr string) error {
	return r.Run(addr)
}
