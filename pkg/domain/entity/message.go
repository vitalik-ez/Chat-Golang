package entity

import "time"

/*
type Message struct {
	ID        uint64
	Data      []byte
	UserID    uint64
	RoomID    uint64
	CreatedAt time.Time
}
*/

type Message struct {
	Room     string    `json:"room"`
	Author   string    `json:"author"`
	Text     string    `json:"text"`
	CreateAt time.Time `json:"time"`
}
