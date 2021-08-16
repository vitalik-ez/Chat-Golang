package entity

import (
	"fmt"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type connection struct {
	ws *websocket.Conn

	send chan []byte
}

func (s subscription) read() {
	fmt.Println("read")
}

func (s *subscription) write() {
	fmt.Println("write")
}

func server() {
	fmt.Println("server")
}
