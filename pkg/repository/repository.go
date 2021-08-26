package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/vitalik-ez/Chat-Golang/pkg/domain/entity"
)

type Authorization interface {
	CreateUser(user entity.User) (int, error)
	GetUser(email, password string) (entity.User, error)
}

type Room interface {
	Create(userId uint64, room entity.Room) (uint64, error)
}

type Message interface {
}

type Repository struct {
	Authorization
	Room
	Message
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		Room:          NewRoomPostgres(db),
	}
}
