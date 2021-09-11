package handler

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/vitalik-ez/Chat-Golang/pkg/domain/entity"
)

// key is name of the room, value is a list save messages
var db = make(map[string][]entity.Message)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

const (
	// Time allowed to write a message to the peer.
	writeWait = 10 * time.Second

	// Time allowed to read the next pong message from the peer.
	pongWait = 60 * time.Second

	// Send pings to peer with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10

	// Maximum message size allowed from peer.
	maxMessageSize = 512

	join = "join"
)

// connection is an middleman between the websocket connection and the hub.
type connection struct {
	// The websocket connection.
	ws *websocket.Conn

	// Buffered channel of outbound messages.
	send chan entity.Message
}

// readPump pumps messages from the websocket connection to the hub.
func (s session) readPump() {
	c := s.conn
	defer func() {
		Hb.leave <- s
		c.ws.Close()
	}()
	c.ws.SetReadLimit(maxMessageSize)
	c.ws.SetReadDeadline(time.Now().Add(pongWait))
	c.ws.SetPongHandler(func(string) error { c.ws.SetReadDeadline(time.Now().Add(pongWait)); return nil })
	for {
		message := entity.Message{}
		err := c.ws.ReadJSON(&message)
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway) {
				log.Printf("error: %v", err)
			}
			break
		}
		fmt.Println("message do broadcast", message)
		Hb.broadcast <- message
	}
}

// write writes a message with the given message type and payload.
func (c *connection) write(mt int, payload []byte) error {
	c.ws.SetWriteDeadline(time.Now().Add(writeWait))
	return c.ws.WriteMessage(mt, payload)
}

// writePump pumps messages from the hub to the websocket connection.
func (s *session) writePump() {
	c := s.conn
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		c.ws.Close()
	}()
	for {
		select {
		case message, ok := <-c.send:
			if !ok {
				c.write(websocket.CloseMessage, []byte{})
				return
			}
			if err := c.ws.WriteJSON(message); err != nil {
				return
			}
		case <-ticker.C:
			if err := c.write(websocket.PingMessage, []byte{}); err != nil {
				return
			}
		}
	}
}

type joinToRoom struct {
	NameRoom string `json:"nameRoom"`
	Username string `json:"username"`
}

type HubCommand struct {
	Command string `json:"command"`
	Data    string `json:"data"`
	Author  string `json:"author"`
}

func (h *Handler) chatRoomWS(c *gin.Context) {

	upgrader.CheckOrigin = func(r *http.Request) bool { return true }

	ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println(err.Error())
		return
	}

	log.Println("Client Successfuly Connected...")

	conn := &connection{send: make(chan entity.Message), ws: ws}
	s := session{conn: conn}

	// Send list exist room
	var existRoom []string
	for key, _ := range db {
		existRoom = append(existRoom, key)
	}
	ws.WriteJSON(existRoom)

	hubCommand := &HubCommand{}
	if err = ws.ReadJSON(&hubCommand); err != nil {
		log.Println(err.Error())
		return
	}
	fmt.Println("Server command: ", hubCommand)

	if hubCommand.Command == join {
		s.commands = hubCommand
		Hb.join <- s
	}
	//Hb.join <- s
	go s.writePump()
	go s.readPump()
}
