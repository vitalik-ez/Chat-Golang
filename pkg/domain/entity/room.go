package entity

type Room struct {
	ID   uint64
	Name string `json:"roomName" binding:"required"`
}
