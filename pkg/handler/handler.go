package handler

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/vitalik-ez/Chat-Golang/pkg/service"
)

type Handler struct {
	services *service.Service
	hub      *hub
}

func NewHandler(services *service.Service, hub *hub) *Handler {
	return &Handler{services: services, hub: hub}
}

func (h *Handler) InitRoutes() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/status-server", func(rw http.ResponseWriter, r *http.Request) {
		rw.WriteHeader(http.StatusOK)
	})

	router.HandleFunc("/api/room/ws/", h.chatRoomWS)

	return router
}
