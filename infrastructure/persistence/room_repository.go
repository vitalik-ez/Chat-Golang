package persistence

import (
	"fmt"

	"github.com/vitalik-ez/Chat-Golang/domain/entity"
)

type RoomRepo struct {
	Db string
}

func (r *RoomRepo) CreateRoom(room *entity.Room) {
	fmt.Println("CreateRoom")

}

func (r *RoomRepo) GetUsersByRoomId(id uint64) {
	fmt.Println("GetUsersByRoomId")
}
