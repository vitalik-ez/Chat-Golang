package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/vitalik-ez/Chat-Golang/pkg/domain/entity"
)

type Room interface {
	Create(room string) error
}

type Message interface {
	Create(message entity.Message) error
}

type Repository struct {
	Room
	Message
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Room:    NewRoomPostgres(db),
		Message: NewMessagePostgres(db),
	}
}
