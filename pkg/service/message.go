package service

import (
	"github.com/vitalik-ez/Chat-Golang/pkg/domain/entity"
	"github.com/vitalik-ez/Chat-Golang/pkg/repository"
)

type MessageService struct {
	repo repository.Message
}

func NewMessageService(repo repository.Message) *MessageService {
	return &MessageService{repo: repo}
}

func (s *MessageService) Create(message entity.Message) error {
	return s.repo.Create(message)
}
