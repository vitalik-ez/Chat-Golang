package service

import (
	"github.com/vitalik-ez/Chat-Golang/pkg/domain/entity"
	"github.com/vitalik-ez/Chat-Golang/pkg/repository"
)

type Authorization interface {
	CreateUser(user entity.User) (int, error)
	GenerateToken(email, password string) (string, error)
	ParseToken(accessToken string) (uint64, error)
}

type Room interface {
	Create(userId uint64, room entity.Room) (uint64, error)
}

type Message interface {
}

type Service struct {
	Authorization
	Room
	Message
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		Room:          NewRoomService(repos.Room),
	}
}
