package handler

import (
	"fmt"
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

//var db = make(map[string][]*entity.Message)

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
			service.Room.Create(s.Client.Room)
			Hb.Join <- s
		case broadcastCommand:
			message := entity.NewMessage(s.Client.Room, s.Client.UserName, s.Client.Message)
			service.Message.Create(*message)
			Hb.Broadcast <- s
		case leaveCommand:
			break Loop //Hb.Leave <- s
		default:
			fmt.Println("Incorrect coomand !!!")
		}
	}
}

func (s *Session) write(mt int, payload []byte) error {
	return s.WS.WriteMessage(mt, payload)
}

func (s *Session) writePump(Hb *hub) {
	//ticker := time.NewTicker(pingPeriod)
	defer func() {
		//ticker.Stop()
		s.WS.Close()
	}()

	for {
		select {
		case message, ok := <-s.Send:
			if !ok {
				fmt.Println("close message", message.Text, ok)
				s.write(websocket.CloseMessage, []byte{})
				return
			}
			if err := s.WS.WriteJSON(message); err != nil {
				return
			}
			/*case <-ticker.C:
			if err := c.write(websocket.PingMessage, []byte{}); err != nil {
				return
			}*/
		}
	}
}
