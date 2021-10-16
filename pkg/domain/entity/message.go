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
	UserName string    `json:"userName"`
	Text     string    `json:"text"`
	CreateAt time.Time `json:"time"`
}

func NewMessage(room string, userName string, text string) *Message {
	return &Message{
		Room:     room,
		UserName: userName,
		Text:     text,
		CreateAt: time.Now(),
	}
}
