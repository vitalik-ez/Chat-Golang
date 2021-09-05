package entity

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
	/*Email    string `json:"email"`
	Username string `json:"username"`*/
	Text string `json:"text"`
}
