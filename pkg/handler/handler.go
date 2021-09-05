package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/vitalik-ez/Chat-Golang/pkg/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {

	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}

	api := router.Group("/api", h.userIdentity)
	{
		room := api.Group("/room")
		{
			room.GET("/", h.getAllRooms)
			room.POST("/", h.createRoom)
			room.GET("/:roomId", h.chatRoom)
			room.GET("/ws/:roomId", h.chatRoomWS)
		}
	}

	return router
}
