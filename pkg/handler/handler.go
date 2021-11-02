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

func (h *Handler) InitRoutes(hb *hub) *gin.Engine {

	router := gin.New()

	router.GET("status-server", h.getStatusServer)
	{
		api := router.Group("/api") //  ,h.userIdentity
		room := api.Group("/room")
		{
			room.GET("/ws/", func(c *gin.Context) {
				h.chatRoomWS(hb, c)
			})
		}
	}

	return router
}
