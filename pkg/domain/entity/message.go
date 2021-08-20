package entity

import "time"

type Message struct {
	ID        uint64
	Data      []byte
	UserID    uint64
	RoomID    uint64
	CreatedAt time.Time
}
