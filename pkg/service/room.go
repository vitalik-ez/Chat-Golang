package service

import (
	"github.com/vitalik-ez/Chat-Golang/pkg/domain/entity"
	"github.com/vitalik-ez/Chat-Golang/pkg/repository"
)

type RoomService struct {
	repo repository.Room
}

func NewRoomService(repo repository.Room) *RoomService {
	return &RoomService{repo: repo}
}

func (s *RoomService) Create(userId uint64, room entity.Room) (uint64, error) {
	return s.repo.Create(userId, room)
}
