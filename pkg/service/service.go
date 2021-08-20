package service

import (
	"github.com/vitalik-ez/Chat-Golang/pkg/domain/entity"
	"github.com/vitalik-ez/Chat-Golang/pkg/repository"
)

type Authorization interface {
	CreateUser(user entity.User) (int, error)
}

type Room interface {
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
	}
}