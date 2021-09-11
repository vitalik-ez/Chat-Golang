package entity

type Room struct {
	ID   uint64
	Name string `json:"name" binding:"required"`
	//FounderId uint64
	//Users []*User
}
