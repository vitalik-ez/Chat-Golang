package handler

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/vitalik-ez/Chat-Golang/pkg/domain/entity"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func (h *Handler) chatRoomWS(w http.ResponseWriter, r *http.Request) {

	upgrader.CheckOrigin = func(r *http.Request) bool { return true }

	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err.Error())
		return
	}

	s := NewSession(make(chan entity.Message), ws)

	go s.writePump(h.hub)
	go s.readPump(h.hub, h.services)
}
