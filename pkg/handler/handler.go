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
		auth.POST("/sign-up", h.signIn)
		auth.POST("/sign-in", h.signUp)
	}

	api := router.Group("/api")
	{
		room := api.Group("/room")
		{
			room.GET("/", h.getAllRooms)
			room.POST("/", h.createRoom)
		}
	}

	return router
}
