package handler

import (
	"fmt"

	"github.com/vitalik-ez/Chat-Golang/pkg/domain/entity"
)

/*
type message struct {
	data []byte
	room string
}*/

type session struct {
	conn     *connection
	commands *HubCommand

	//room string
}

type hub struct {
	rooms     map[string]map[*connection]bool
	broadcast chan entity.Message
	join      chan session
	leave     chan session
}

var Hb = hub{
	broadcast: make(chan entity.Message),
	join:      make(chan session),
	leave:     make(chan session),
	rooms:     make(map[string]map[*connection]bool),
}

func (h *hub) Run() {
	for {
		select {
		case s := <-h.join:
			connections := h.rooms[s.commands.Data]
			if connections == nil {
				connections = make(map[*connection]bool)
				h.rooms[s.commands.Data] = connections
			}
			h.rooms[s.commands.Data][s.conn] = true
			fmt.Println("check ", h.rooms)
		/*case s := <-h.leave:
		connections := h.rooms[s.room]
		if connections != nil {
			if _, ok := connections[s.conn]; ok {
				delete(connections, s.conn)
				close(s.conn.send)
				if len(connections) == 0 {
					delete(h.rooms, s.room)
				}
			}
		}*/
		case m := <-h.broadcast:
			connections := h.rooms[m.Room]
			fmt.Println("Room", m.Room, "connections", connections)
			db[m.Room] = append(db[m.Room], m)
			fmt.Println("broadcast", m)
			for c := range connections {
				fmt.Println("c", c)
				select {
				case c.send <- m:
				default:
					close(c.send)
					delete(connections, c)
					if len(connections) == 0 {
						delete(h.rooms, m.Room)
					}
				}
			}
		}
	}
}
