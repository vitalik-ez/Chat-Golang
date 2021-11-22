package handler

import (
	"log"

	"github.com/gorilla/websocket"
	"github.com/vitalik-ez/Chat-Golang/pkg/domain/entity"
	"github.com/vitalik-ez/Chat-Golang/pkg/service"
)

const (
	maxMessageSize   = 512
	joinCommand      = "join"
	broadcastCommand = "broadcast"
	leaveCommand     = "leave"
)

type HubCommand struct {
	Command  string `json:"command"`
	Message  string `json:"message"`
	UserName string `json:"userName"`
	Room     string `json:"room"`
}

type Session struct {
	WS     *websocket.Conn
	Send   chan entity.Message
	Client HubCommand
}

func NewSession(send chan entity.Message, ws *websocket.Conn) *Session {
	return &Session{
		Send: send,
		WS:   ws,
	}
}

func (s Session) readPump(Hb *hub, service *service.Service) {
	defer func() {
		Hb.Leave <- s
		s.WS.Close()
	}()
	s.WS.SetReadLimit(maxMessageSize)
Loop:
	for {
		err := s.WS.ReadJSON(&s.Client)
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			break
		}
		switch s.Client.Command {
		case joinCommand:
			if err = service.Room.Create(s.Client.Room); err != nil {
				log.Println("Error writing data to database:", err.Error())
			}
			Hb.Join <- s
		case broadcastCommand:
			message := entity.NewMessage(s.Client.Room, s.Client.UserName, s.Client.Message)
			if err = service.Message.Create(*message); err != nil {
				log.Println("Error writing data to database:", err.Error())
			}
			Hb.Broadcast <- s
		case leaveCommand:
			break Loop //Hb.Leave <- s
		default:
			log.Println("Incorrect coomand")
		}
	}
}

func (s *Session) write(mt int, payload []byte) error {
	return s.WS.WriteMessage(mt, payload)
}

func (s *Session) writePump(Hb *hub) {
	defer func() {
		s.WS.Close()
	}()

	for {
		select {
		case message, ok := <-s.Send:
			if !ok {
				log.Println("close message", message.Text, ok)
				s.write(websocket.CloseMessage, []byte{})
				return
			}
			if err := s.WS.WriteJSON(message); err != nil {
				return
			}
		}
	}
}
