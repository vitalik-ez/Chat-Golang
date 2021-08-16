package persistence

import "github.com/vitalik-ez/Chat-Golang/domain/repository"

type Repositories struct {
	User    repository.UserRepository
	Message repository.MessageRepository
	Room    repository.RoomRepository
	db      string
}

func NewRepositories() {

}

//closes the  database connection
func (s *Repositories) Close() {
	//return s.db.Close()
}
