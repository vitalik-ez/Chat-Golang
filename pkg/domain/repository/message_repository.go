package repository

import "github.com/vitalik-ez/Chat-Golang/domain/entity"

type MessageRepository interface {
	SaveMessage(*entity.Message)
	GetAllMessageByRoomId(id uint64)
}
