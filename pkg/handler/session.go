package handler

import (
	"fmt"
	"log"

	"github.com/gorilla/websocket"
	"github.com/vitalik-ez/Chat-Golang/pkg/domain/entity"
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

var db = make(map[string][]*entity.Message)

func (s Session) readPump(Hb *hub) {
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
			Hb.Join <- s
		case broadcastCommand:
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
