package service

import (
	"github.com/vitalik-ez/Chat-Golang/pkg/domain/entity"
	"github.com/vitalik-ez/Chat-Golang/pkg/repository"
)

//go:generate mockgen -source=service.go -destination=mocks/mock.go

type Room interface {
	Create(room string) error
}

type Message interface {
	Create(message entity.Message) error
}

type Service struct {
	Room
	Message
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Room:    NewRoomService(repos.Room),
		Message: NewMessageService(repos.Message),
	}
}
