package handler

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/vitalik-ez/Chat-Golang/pkg/domain/entity"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func (h *Handler) chatRoomWS(hb *hub, c *gin.Context) {

	upgrader.CheckOrigin = func(r *http.Request) bool { return true }

	ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println(err.Error())
		return
	}

	log.Println("Client Successfuly Connected...")

	s := Session{Send: make(chan entity.Message), WS: ws}

	go s.writePump(hb)
	go s.readPump(hb)
}
