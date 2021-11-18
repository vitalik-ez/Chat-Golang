package handler

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/vitalik-ez/Chat-Golang/pkg/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes(hb *hub) *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/status-server", func(rw http.ResponseWriter, r *http.Request) {
		rw.WriteHeader(http.StatusOK)
	})

	router.HandleFunc("/api/room/ws/", func(rw http.ResponseWriter, r *http.Request) {
		h.chatRoomWS(hb, rw, r)
	})

	return router
}
