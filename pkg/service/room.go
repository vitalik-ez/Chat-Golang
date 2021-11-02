package service

import (
	"github.com/vitalik-ez/Chat-Golang/pkg/repository"
)

type RoomService struct {
	repo repository.Room
}

func NewRoomService(repo repository.Room) *RoomService {
	return &RoomService{repo: repo}
}

func (s *RoomService) Create(room string) error {
	return s.repo.Create(room)
}
