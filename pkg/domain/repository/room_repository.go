package repository

import "github.com/vitalik-ez/Chat-Golang/domain/entity"

type RoomRepository interface {
	CreateRoom(*entity.Room)
	GetUsersByRoomId(id uint64)
}
