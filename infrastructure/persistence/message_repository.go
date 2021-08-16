package persistence

import (
	"fmt"

	"github.com/vitalik-ez/Chat-Golang/domain/entity"
)

type MessageRepo struct {
	Db string
}

func (m *MessageRepo) SaveMessage(user *entity.Message) {
	fmt.Println("Save Message")

}

func (m *MessageRepo) GetAllMessageByRoomId(id uint64) {
	fmt.Println("GetAllMessageByRoomId")
}
