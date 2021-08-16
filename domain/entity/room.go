package entity

type Room struct {
	ID    uint64
	Name  string
	Users []*User
}
