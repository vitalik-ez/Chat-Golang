package handler

import (
	"fmt"

	"github.com/vitalik-ez/Chat-Golang/pkg/domain/entity"
)

type hub struct {
	Rooms     map[string]map[*Session]bool
	Broadcast chan Session
	Join      chan Session
	Leave     chan Session
}

func NewHub() *hub {
	return &hub{
		Rooms:     make(map[string]map[*Session]bool),
		Broadcast: make(chan Session),
		Join:      make(chan Session),
		Leave:     make(chan Session),
	}
}

func (h *hub) Run() {
	for {
		select {
		case s := <-h.Join:
			connections := h.Rooms[s.Client.Room]
			if connections == nil {
				connections = make(map[*Session]bool)
				h.Rooms[s.Client.Room] = connections
			}
			h.Rooms[s.Client.Room][&s] = true
		case s := <-h.Leave:
			fmt.Println("Leave client", s.Client)
			connections := h.Rooms[s.Client.Room]
			if connections != nil {
				if _, ok := connections[&s]; ok {
					delete(connections, &s)
					close(s.Send)
					if len(connections) == 0 {
						delete(h.Rooms, s.Client.Room)
					}
				}
			}
		case m := <-h.Broadcast:
			message := entity.NewMessage(m.Client.Room, m.Client.UserName, m.Client.Message)
			connections := h.Rooms[m.Client.Room]
			// add to tmp db
			db[m.Client.Room] = append(db[m.Client.Room], message)
			for c := range connections {
				if c.WS != m.WS {
					select {
					case c.Send <- *message:

					default:
						close(c.Send)
						delete(connections, c)
						if len(connections) == 0 {
							delete(h.Rooms, m.Client.Room)
						}
					}
				}
			}
		}
	}
}
